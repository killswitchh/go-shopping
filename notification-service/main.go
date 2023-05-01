package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/api", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello from Notification service",
		})
	})

	router.Run(":3001")
}