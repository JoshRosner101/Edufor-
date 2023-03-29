package main

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

// set up router
func SetUpRouter() *gin.Engine {
	connectDB("./mock.db")
	router := gin.Default()
	return router
}

func TestGetThreads(t *testing.T) {
	r := SetUpRouter()
	r.GET("/backend/threads", getThreads)
	req, _ := http.NewRequest("GET", "/backend/threads", nil)
	// create response recorder
	w := httptest.NewRecorder()
	// perform request
	r.ServeHTTP(w, req)

	/*
		var threads []Thread
		json.Unmarshal(w.Body.Bytes(), &threads)
	*/

	// check if response was as expected
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteThreadByID(t *testing.T) {
	db.Exec("DELETE FROM thread")
	db.Exec("DELETE FROM reply")
	db.Exec("DELETE FROM user")
	r := SetUpRouter()
	r.POST("/backend/threads", postThreads)
	thread := Thread{
		ID:       1,
		Username: "poster27",
		Title:    "I need some help with a project",
		Body:     "I've been working on a piece of software and could use a helping hand.",
		Time:     "2/7/2023, 1:43:27 PM",
		Replies:  nil,
	}

	jsonValue, _ := json.Marshal(thread)
	req, _ := http.NewRequest("POST", "/backend/threads", bytes.NewBuffer(jsonValue))
	// create response recorder
	w := httptest.NewRecorder()
	// perform request
	r.ServeHTTP(w, req)

	r.DELETE("/backend/threads/1", deleteThreadByID)
	req2, _ := http.NewRequest("DELETE", "/backend/threads/1", nil)
	// perform request
	w2 := httptest.NewRecorder()
	r.ServeHTTP(w2, req2)
	assert.Equal(t, http.StatusOK, w2.Code)
}

func TestPutThread(t *testing.T) {
	db.Exec("DELETE FROM thread")
	db.Exec("DELETE FROM reply")
	db.Exec("DELETE FROM user")
	r := SetUpRouter()
	r.POST("/backend/threads", postThreads)
	thread := Thread{
		ID:       1,
		Username: "poster27",
		Title:    "I need some help with a project",
		Body:     "I've been working on a piece of software and could use a helping hand.",
		Time:     "2/7/2023, 1:43:27 PM",
		Replies:  nil,
	}

	jsonValue, _ := json.Marshal(thread)
	req, _ := http.NewRequest("POST", "/backend/threads", bytes.NewBuffer(jsonValue))
	// create response recorder
	w := httptest.NewRecorder()
	// perform request
	r.ServeHTTP(w, req)

	thread2 := Thread{
		ID:       3,
		Username: "randomUser1",
		Title:    "Anyone want to work on a project?",
		Body:     "I made a better ChatGPT",
		Time:     "3/21/2023, 2:20:28 PM",
		Replies:  nil,
	}
	r.PUT("/backend/threads/1", putThread)
	jsonValue2, _ := json.Marshal(thread2)
	req2, _ := http.NewRequest("PUT", "/backend/threads/1", bytes.NewBuffer(jsonValue2))
	// perform request
	w2 := httptest.NewRecorder()
	r.ServeHTTP(w2, req2)
	assert.Equal(t, http.StatusOK, w2.Code)
}

/*
func TestGetThreadsByID(t *testing.T) {
	r := SetUpRouter()
	r.GET("/backend/threads/1", getThreadByID)
	req, _ := http.NewRequest("GET", "/backend/threads/1", nil)
	// create response recorder
	w := httptest.NewRecorder()
	// perform request
	r.ServeHTTP(w, req)

	var threads []Thread
	json.Unmarshal(w.Body.Bytes(), &threads)

	// check if response was as expected
	assert.Equal(t, http.StatusOK, w.Code)
}
*/

// clearing database
func TestClearDB(t *testing.T) {
	db.Exec("DELETE FROM thread")
	db.Exec("DELETE FROM reply")
	db.Exec("DELETE FROM user")
}

func TestPostThreads(t *testing.T) {
	db.Exec("DELETE FROM thread")
	db.Exec("DELETE FROM reply")
	db.Exec("DELETE FROM user")
	r := SetUpRouter()
	r.POST("/backend/threads", postThreads)
	thread := Thread{
		ID:       1,
		Username: "poster27",
		Title:    "I need some help with a project",
		Body:     "I've been working on a piece of software and could use a helping hand.",
		Time:     "2/7/2023, 1:43:27 PM",
		Replies:  nil,
	}

	jsonValue, _ := json.Marshal(thread)
	req, _ := http.NewRequest("POST", "/backend/threads", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestPostReply(t *testing.T) {
	db.Exec("DELETE FROM thread")
	db.Exec("DELETE FROM reply")
	db.Exec("DELETE FROM user")
	r := SetUpRouter()
	r.POST("/backend/threads/1", postReply)
	reply := Reply{
		ReplyID:   1,
		Username:  "poster27",
		Body:      "I've been working on a piece of software and could use a helping hand.",
		Time:      "2/7/2023, 1:43:27 PM",
		ReplyPost: 1,
	}

	jsonValue, _ := json.Marshal(reply)
	req, _ := http.NewRequest("POST", "/backend/threads/1", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestAllThreads(t *testing.T) {
	thread := []Thread{{
		ID:       1,
		Username: "poster27",
		Title:    "I need some help with a project",
		Body:     "I've been working on a piece of software and could use a helping hand.",
		Time:     "2/7/2023, 1:43:27 PM",
		Replies:  nil,
	}}
	actual, _ := allThreads()
	expected := thread

	assert.Equal(t, expected, actual)
}

func TestThreadByID(t *testing.T) {
	thread := Thread{
		ID:       1,
		Username: "poster27",
		Title:    "I need some help with a project",
		Body:     "I've been working on a piece of software and could use a helping hand.",
		Time:     "2/7/2023, 1:43:27 PM",
		Replies:  nil,
	}
	actual, _ := threadByID(1)
	expected := thread

	assert.Equal(t, expected, actual)
}

func TestAddReply(t *testing.T) {
	rep := Reply{
		ReplyID:   1,
		Username:  "poster27",
		Body:      "I've been working on a piece of software and could use a helping hand.",
		Time:      "2/7/2023, 1:43:27 PM",
		ReplyPost: 1,
	}
	addReply(rep)
	rows, _ := db.Query("SELECT * FROM reply WHERE replyid = 1")

	i := 0
	for rows.Next() {
		i++
	}
	assert.NotEqual(t, 0, i)
}

/*
func TestReplyByPostID(t *testing.T) {
}
*/

func TestReplaceThread(t *testing.T) {
	db.Exec("DELETE FROM thread")
	db.Exec("DELETE FROM reply")
	db.Exec("DELETE FROM user")

	thread := Thread{
		ID:       1,
		Username: "poster27",
		Title:    "I need some help with a project",
		Body:     "I've been working on a piece of software and could use a helping hand.",
		Time:     "2/7/2023, 1:43:27 PM",
		Replies:  nil,
	}

	addThread(thread)

	expected := Thread{
		ID:       1,
		Username: "poster27",
		Title:    "I need some help with a project",
		Body:     "I have finished on this piece of software. Thanks for the help! (edited)",
		Time:     "2/7/2023, 1:43:27 PM",
		Replies:  nil,
	}

	replaceThread(1, expected)
	actual, _ := threadByID(1)
	assert.Equal(t, expected, actual)
}

func TestRegister(t *testing.T) {
	r := SetUpRouter()
	r.POST("/users/register", register)
	//Test register with new user
	user := User{
		UserID:   1,
		Name:     "George",
		Password: "TestPassword",
	}

	jsonValue, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/users/register", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)

	//Test register with already existing user (should be an error)

	req, _ = http.NewRequest("POST", "/users/register", bytes.NewBuffer(jsonValue))

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}
func TestLogin(t *testing.T) {
	r := SetUpRouter()
	r.POST("/users/login", login)
	//Login with wrong user
	user := User{
		UserID:   1,
		Name:     "George2",
		Password: "TestPassword",
	}

	jsonValue, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/users/login", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)

	//Login with wrong password

	user = User{
		UserID:   1,
		Name:     "George",
		Password: "WrongPassword",
	}

	jsonValue, _ = json.Marshal(user)

	req, _ = http.NewRequest("POST", "/users/login", bytes.NewBuffer(jsonValue))

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)

	//Login correctly

	user = User{
		UserID:   1,
		Name:     "George",
		Password: "TestPassword",
	}

	jsonValue, _ = json.Marshal(user)

	req, _ = http.NewRequest("POST", "/users/login", bytes.NewBuffer(jsonValue))

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUserByName(t *testing.T) {
	user := User{
		UserID:   1,
		Name:     "George",
		Password: "",
	}
	actual, _ := userByName("George")
	//we don't need to compare the password in this
	actual.Password = ""
	expected := user

	assert.Equal(t, actual, expected)
}

func TestUserByID(t *testing.T) {
	user := User{
		UserID:   1,
		Name:     "George",
		Password: "",
	}
	actual, _ := userByID(1)
	//we don't need to compare the password in this
	actual.Password = ""
	expected := user

	assert.Equal(t, actual, expected)
}

//These two tests require cookies to function, and I have been unable to find a reliable way to test cookies with the current working code without rewriting part of it
//Instead of using the golang testing library for these functions, these will be tested with postman since it supports cookies.

/*
func TestCurrentUser(t *testing.T) {
	r := SetUpRouter()
	r.POST("/users/login", login)
	r.GET("/users/user", currentUser)
	//Test when Logged in

	user := User{
		UserID:      			 1,
		Name: 			  "George",
		Password:   "TestPassword",
	}

	jsonValue, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/users/login", bytes.NewBuffer(jsonValue))

	req, _ = http.NewRequest("GET", "/users/user", nil)
	w := httptest.NewRecorder()
	// perform request
	r.ServeHTTP(w, req)

	// check if response was as expected
	assert.Equal(t, http.StatusOK, w.Code)

}
func TestLogout(t *testing.T) {
	r := SetUpRouter()
	r.POST("/users/logout", logout)
	//Test log out

	//logout doesnt care about the json value

	user := User{
		UserID:      			 1,
		Name: 			  "George",
		Password:   "TestPassword",
	}

	jsonValue, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/users/login", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	//Test currentUser when logged out
}
*/
