package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  200,
			"message": "The server is ready.",
			"time":    time.Now(),
		})
	})

	router.NoRoute(func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  404,
			"message": "Not Found",
			"time":    time.Now(),
		})
	})

	router.Run(":80")
}
