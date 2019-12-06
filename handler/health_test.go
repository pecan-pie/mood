package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/pecan-pie/mood/handler"
	"github.com/stretchr/testify/assert"
)

func TestHealthHandler(t *testing.T) {

	router := gin.Default()
	router.GET("/health", handler.Health)

	// Define the request
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Define the request recorder
	rr := httptest.NewRecorder()

	// Mock Server
	router.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)
	assert.Equal(t, "ok", rr.Body.String())
}
