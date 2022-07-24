package main

import (
	"net/http"

	"learn-gin/controllers"
	"learn-gin/models" // 导入项目下的包格式为  workspace/packname

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default() // router with default middleware installed

	// 一组简单但有效的 CRUD RESTful API 路由！
	models.ConnectDatabase()                      // new!
	router.POST("/posts", controllers.CreatePost) // here!
	router.GET("/posts", controllers.FindPosts)
	router.GET("/posts/:id", controllers.FindPost) // here!
	router.PATCH("/posts/:id", controllers.UpdatePost)
	router.DELETE("/posts/:id", controllers.DeletePost)

	// index route
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World lenrn gin now!",
		})
	})
	// run the server
	router.Run()
}
