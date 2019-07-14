package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Good Bye: " + time.Now().Format(time.RFC1123),
		})
	})

	router.Run(":8080")
}
