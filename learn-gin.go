package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default() // router with default middleware installed
	// index route
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World lenrn gin now!",
		})
	})
	// run the server
	router.Run()
}
