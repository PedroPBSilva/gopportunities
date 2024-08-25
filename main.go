package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main ()  {
	fmt.Println("Starting API Rest!")
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run(":8080") 
}
