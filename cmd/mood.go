package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pecan-pie/mood/handler"
)

func main() {
	r := gin.Default()
	r.GET("/health", handler.Health)
	r.Run(":8080")
}
