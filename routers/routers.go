package routers

import (
	"goblog/controller"

	"github.com/gin-gonic/gin"
)

func Start() {
	e := gin.Default()
	e.LoadHTMLGlob("templates/*")
	e.Static("/assets", "./assets")

	e.POST("/register", controller.Register)
	e.GET("/register", controller.GoRegister)

	e.POST("/login", controller.Login)
	e.GET("/login", controller.GoLogin)

	e.GET("/post_index", controller.GetPostIndex)
	e.POST("/post", controller.AddPost)
	e.GET("/post", controller.GoAddPost)
	e.GET("/detail", controller.PostDetail)

	e.GET("/", controller.Index)
	e.Run()
}
