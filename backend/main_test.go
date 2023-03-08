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
}

func TestPostThreads(t *testing.T) {
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

	assert.Equal(t, actual, expected)
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

	assert.Equal(t, actual, expected)
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
