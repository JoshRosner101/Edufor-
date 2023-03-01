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
	connectDB()
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

	var threads []Thread
	json.Unmarshal(w.Body.Bytes(), &threads)

	// check if response was as expected
	assert.Equal(t, http.StatusOK, w.Code)
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

/*
func TestAllThreads(t *testing.T) {
}
func TestThreadsByID(t *testing.T) {
}
func TestThreadsByUsername(t *testing.T) {
}
func TestAddReply(t *testing.T) {
}
func TestReplyByPostID(t *testing.T) {
}
*/
