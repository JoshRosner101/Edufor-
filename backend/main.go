package main

import (
    "net/http"
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "strconv"
    "database/sql"
    "fmt"
    "log"
    "golang.org/x/crypto/bcrypt"
    "github.com/golang-jwt/jwt/v5"
    "time"
    _ "modernc.org/sqlite"
)

var db *sql.DB
const secretKey = "secret"

type User struct {
    UserID      int64 `json:"userid"`
    Name        string `json:"name"`
    Password    string `json:"-"`
}

type Reply struct {
    //ID of each unique reply
    ReplyID     int64  `json:"replyid"`
    //Username of the one who posted it
    Username string `json:"username"`
    //The main text in a post
    Body string `json:"body"`
    //The time when the post was made
    Time string `json:"time"`
    //ID of the post this reply references
    ReplyPost     int64  `json:"replypost"`
}

// thread represents data about each individual post.
type Thread struct {
    //ID of each unique post
    ID     int64  `json:"id"`
    //Username of the one who posted it
    Username string `json:"username"`
    //Title of a post
    Title  string  `json:"title"`
    //The main text in a post
    Body string `json:"body"`
    //The time when the post was made
    Time string `json:"time"`
    //This stores all of the Replies for each thread
    Replies []Reply `json:"replies"`
}

func main() {
    connectDB("./threads.db")
    router := gin.Default()

    router.Use(cors.New(cors.Config{
        AllowCredentials: true,
        AllowOrigins:     []string{"https://http://localhost:4200"},
    }))

    backend := router.Group("/backend")
    {
        backend.GET("/threads", getThreads)
        backend.POST("/threads", postThreads)

        //Currently disabled because :username conflicts with :id
        //backend.GET("/threads/:username", getThreadsByUsername)

        backend.GET("/threads/:id", getThreadByID)
        backend.POST("/threads/:id", postReply)
        //backend.PUT("/threads/:id", [put an update funciton here])
        //backend.DELETE("/threads/:id", [put a delete function here])

        //Registers user with encrypted password
        backend.POST("/users/register", register)
        //Logs in
        backend.POST("/users/login", login)

        //Checks whether you're logged in
        backend.GET("/users/user", currentUser)
        //Logs out
        backend.POST("/users/logout", logout)
    }

    router.Run("0.0.0.0:8080")
    
}

//This function opens the database
func connectDB(pathname string) error {
    DB, err := sql.Open("sqlite",pathname)
    if err != nil {
        log.Fatal(err)
    }
    db = DB
    return nil
}


// getThreads responds with the list of all threads as JSON.
func getThreads(c *gin.Context) {
    threads, err := allThreads()
    if err != nil {
        log.Fatal(err)
    }
    c.IndentedJSON(http.StatusOK, threads)
}

// getThreadByID locates the thread whose ID value matches the id
// parameter sent by the client, then returns that thread as a response.
func getThreadByID(c *gin.Context) {
    id, err := strconv.ParseInt(c.Param("id"),10,64)

    //This gets the main post of the thread
    post, err := threadByID(id)
    if err != nil {
        //log.Fatal(err)
        c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return;
    }

    //This gets the replies for the thread and embeds them into the thread as an array.
    postReplies, err := replyByPostID(id)
    if err != nil {
        log.Fatal(err)
    }
    post.Replies = postReplies

    c.IndentedJSON(http.StatusOK, post)
}

//This uses the gin router to post all threads by a certain user
func getThreadsByUsername(c *gin.Context) {
    username := c.Param("username")

    threads, err := threadsByUsername(username)
    if err != nil {
        log.Fatal(err)
    }
    c.IndentedJSON(http.StatusOK, threads)
}

// postThreads adds a thread from JSON received in the request body.
func postThreads(c *gin.Context) {
    var newThread Thread

    // Call BindJSON to bind the received JSON to
    // newThread.
    if err := c.BindJSON(&newThread); err != nil {
        return
    }

    threadID, err := addThread(newThread)
    if err != nil {
        log.Fatal(err)
    }
    //The id of the new thread is printed to console; may be used for something else in the future.
    fmt.Printf("ID of added thread: %v\n", threadID)
    c.IndentedJSON(http.StatusCreated, newThread)
}

// postThreads adds a thread from JSON received in the request body.
func postReply(c *gin.Context) {
    var newReply Reply

    // Call BindJSON to bind the received JSON to
    // newThread.
    if err := c.BindJSON(&newReply); err != nil {
        return
    }

    replyID, err := addReply(newReply)
    if err != nil {
        log.Fatal(err)
    }
    //The id of the new thread is printed to console; may be used for something else in the future.
    fmt.Printf("ID of added reply: %v\n", replyID)
    c.IndentedJSON(http.StatusCreated, newReply)
}


//Functions that interact with database down here

//This function adds a new thread to the database (passes in a thread object and modifies it)
//Returns the id of the post (may be used)
func addThread(post Thread) (int64, error) {
    result, err := db.Exec("INSERT INTO thread (username, title, body, time) VALUES (?, ?, ?, ?)", post.Username, post.Title, post.Body, post.Time)
    if err != nil {
        return 0, fmt.Errorf("addThread: %v", err)
    }
    id, err := result.LastInsertId()
    if err != nil {
        return 0, fmt.Errorf("addThread: %v", err)
    }
    return id, nil
}

//This function gets all the threads in the database
func allThreads() ([]Thread, error) {
    // A threads slice to hold data from returned rows.
    var threads []Thread

    rows, err := db.Query("SELECT * FROM thread")
    if err != nil {
        return nil, fmt.Errorf("allThreads %v", err)
    }
    defer rows.Close()
    for rows.Next() {
        var post Thread
        if err := rows.Scan(&post.ID, &post.Username, &post.Title, &post.Body, &post.Time); err != nil {
            return nil, fmt.Errorf("allThreads %v", err)
        }
        threads = append(threads, post)
    }
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("allThreads %v", err)
    }
    return threads, nil
}

func threadByID(id int64) (Thread, error) {
    // A post to hold data from the returned row.
    var post Thread

    row := db.QueryRow("SELECT * FROM thread WHERE id = ?", id)
    if err := row.Scan(&post.ID, &post.Username, &post.Title, &post.Body, &post.Time); err != nil {
        if err == sql.ErrNoRows {
            return post, fmt.Errorf("threadByID %d: no such thread", id)
        }
        return post, fmt.Errorf("threadByID %d: %v", id, err)
    }
    return post, nil
}

//This function gets all the threads made by a certain user (may be used for a profile page?)
func threadsByUsername(name string) ([]Thread, error) {
    // A threads slice to hold data from returned rows.
    var threads []Thread

    rows, err := db.Query("SELECT * FROM thread WHERE username = ?", name)
    if err != nil {
        return nil, fmt.Errorf("threadsByUsername %q: %v", name, err)
    }
    defer rows.Close()
    // Loop through rows, using Scan to assign column data to struct fields.
    for rows.Next() {
        var post Thread
        if err := rows.Scan(&post.ID, &post.Username, &post.Title, &post.Body, &post.Time); err != nil {
            return nil, fmt.Errorf("threadsByUsername %q: %v", name, err)
        }
        threads = append(threads, post)
    }
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("threadsByUsername %q: %v", name, err)
    }
    return threads, nil
}

//These functions are for replies

//This function adds a reply into the database.
func addReply(post Reply) (int64, error) {
    result, err := db.Exec("INSERT INTO reply (username, body, time, replypost) VALUES (?, ?, ?, ?)", post.Username, post.Body, post.Time, post.ReplyPost)
    if err != nil {
        return 0, fmt.Errorf("addReply: %v", err)
    }
    id, err := result.LastInsertId()
    if err != nil {
        return 0, fmt.Errorf("addReply: %v", err)
    }
    return id, nil
}

//This function gets all the replies in the database based on the id of the main post.
func replyByPostID(id int64) ([]Reply, error) {
    // A replies slice to hold data from returned rows.
    var replies []Reply

    rows, err := db.Query("SELECT * FROM reply WHERE replypost = ?", id)
    if err != nil {
        return nil, fmt.Errorf("replyByPostID %v", err)
    }
    defer rows.Close()
    for rows.Next() {
        var post Reply
        if err := rows.Scan(&post.ReplyID, &post.Username, &post.Body, &post.Time, &post.ReplyPost); err != nil {
            return nil, fmt.Errorf("replyByPostID %v", err)
        }
        replies = append(replies, post)
    }
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("replyByPostID %v", err)
    }
    return replies, nil
}

//These functions are for logins

func register(c *gin.Context){
    var newUser User

    if err := c.BindJSON(&newUser); err != nil {
        return
    }

    //If username already exists, then there should be an error
    _, err := userByName(newUser.Name)
    if err == nil {
        //log.Fatal(err)
        c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "username already exists"})
        return
    }

    userID, err := addUser(&newUser)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("ID of added user: %v\n", userID)
    c.IndentedJSON(http.StatusCreated, newUser)
}

func login(c *gin.Context){
    var user User

    if err := c.BindJSON(&user); err != nil {
        return
    }

    storedUser, err := userByName(user.Name)
    if err != nil {
        c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "incorrect password"})
        return
    }

    claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
        Issuer: strconv.Itoa(int(storedUser.UserID)),
        ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), //expires after 1 day
    })

    token, err := claims.SignedString([]byte(secretKey))
    if err != nil {
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "could not login"})
        return
    }

    //This sets the cookie
    c.SetCookie("jwt", token, 60*60*24, "/", "localhost", false, true)

    c.IndentedJSON(http.StatusOK, gin.H{"message": "success"})
}

func currentUser(c *gin.Context) {
    cookie, _ := c.Cookie("jwt")

    token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
        return []byte(secretKey), nil
    })
    if err != nil {
        c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "unauthenticated"})
        return
    }
    claims := token.Claims.(*jwt.RegisteredClaims)

    userid, _ := strconv.ParseInt(claims.Issuer, 10, 64)
    user, err := userByID(userid)
    if err != nil {
        c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.IndentedJSON(http.StatusOK, user)
}

func logout(c *gin.Context) {
    //This makes the cookie expire instantly
    c.SetCookie("jwt", "", -1, "/", "localhost", false, true)
    c.IndentedJSON(http.StatusOK, gin.H{"message": "success"})
}

func addUser(post *User) (int64, error) {
    password, _ := bcrypt.GenerateFromPassword([]byte(post.Password),14)
    post.Password = string(password)
    fmt.Printf("Encrypted password: %v\n", post.Password)
    result, err := db.Exec("INSERT INTO user (name, password) VALUES (?, ?)", post.Name, post.Password)
    if err != nil {
        return 0, fmt.Errorf("addUser: %v", err)
    }
    id, err := result.LastInsertId()
    if err != nil {
        return 0, fmt.Errorf("addUser: %v", err)
    }
    post.UserID = id
    return id, nil
}

func userByName(username string) (User, error) {
    var user User

    row := db.QueryRow("SELECT * FROM user WHERE name = ?", username)
    if err := row.Scan(&user.UserID, &user.Name, &user.Password); err != nil {
        if err == sql.ErrNoRows {
            return user, fmt.Errorf("userByName %s: no such user", username)
        }
        return user, fmt.Errorf("userByName %s: %v", username, err)
    }
    return user, nil
}

func userByID(userid int64) (User, error) {
    var user User

    row := db.QueryRow("SELECT * FROM user WHERE userid = ?", userid)
    if err := row.Scan(&user.UserID, &user.Name, &user.Password); err != nil {
        if err == sql.ErrNoRows {
            return user, fmt.Errorf("userByID %d: no such user", userid)
        }
        return user, fmt.Errorf("userByID %d: %v", userid, err)
    }
    return user, nil
}