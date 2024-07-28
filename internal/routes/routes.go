package routes

import (
	"myblog/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	//配置静态文件和模板
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")

	//配置路由
	r.GET("/", handlers.ShowIndexPage)
	r.GET("/posts", handlers.GetPosts)
	r.POST("/posts", handlers.CreatePost)
	r.GET("/posts/new", handlers.ShowNewPostPage)
	r.GET("/posts/:id", handlers.GetPostByID)
	r.PUT("/posts/:id", handlers.UpdatePost)
	r.DELETE("/posts/:id", handlers.DeletePost)

	return r
}
