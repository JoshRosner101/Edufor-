package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

// thread represents data about each individual post.
type thread struct {
    //ID of each unique post
    ID     string  `json:"id"`
    //Username of the one who posted it
    Username string `json:"username"`
    //Title of a post
    Title  string  `json:"title"`
    //The main text in a post
    Body string `json:"body"`
}

// threads slice to seed record thread data.
var threads = []thread{
    {ID: "1", Username: "poster27", Title: "I need some help with a project", Body: "I've been working on a piece of software and could use a helping hand."},
    {ID: "2", Username: "TheRealGogle", Title: "How do I write a database in golang?", Body: "I want to try learning golang but I don't know where to start. I could use some help"},
    {ID: "3", Username: "someRandomName", Title: "Helpful Angular Tips", Body: "I've been working on software for a long time and I have a few helpful pointers as to how you could use angular as the main interface for your next software project."},
}

func main() {
    router := gin.Default()

    backend := router.Group("/backend")
    {
        backend.GET("/threads", getThreads)
        backend.GET("/threads/:id", getThreadByID)
        backend.POST("/threads", postThreads)
    }

    router.Run("0.0.0.0:8080")
}

// getThreads responds with the list of all threads as JSON.
func getThreads(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, threads)
}

// postThreads adds a thread from JSON received in the request body.
func postThreads(c *gin.Context) {
    var newThread thread

    // Call BindJSON to bind the received JSON to
    // newThread.
    if err := c.BindJSON(&newThread); err != nil {
        return
    }

    // Add the new thread to the slice.
    threads = append(threads, newThread)
    c.IndentedJSON(http.StatusCreated, newThread)
}

// getThreadByID locates the thread whose ID value matches the id
// parameter sent by the client, then returns that thread as a response.
func getThreadByID(c *gin.Context) {
    id := c.Param("id")

    // Loop over the list of threads, looking for
    // a thread whose ID value matches the parameter.
    for _, a := range threads {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "thread not found"})
}