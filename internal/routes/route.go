package routes

import (
	"Z02zz/internal/controllers"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func SetupRouter(postController *controllers.PostController, userController *controllers.UserController) *gin.Engine {
	router := gin.Default()

	// 加载 HTML 模板
	router.LoadHTMLGlob(filepath.Join("templates", "*.html"))

	// 设置静态文件目录
	router.Static("/public", "./public")

	// 设置路由
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	api := router.Group("/api")
	{
		// 文章相关路由
		api.GET("/posts", postController.GetPosts)
		api.GET("/posts/:id", postController.GetPost)
		api.GET("/posts/new", postController.NewPostForm)
		api.POST("/posts", postController.CreatePost)
		api.PUT("/posts/:id", postController.UpdatePost)
		api.DELETE("/posts/:id", postController.DeletePost)
	}

	user := router.Group("/user")
	{
		// 用户相关路由
		user.POST("/register", userController.Register)             // 用户注册
		user.POST("/login", userController.Login)                   // 用户登录
		user.PUT("/user/profile", userController.UpdateUserProfile) // 更新用户信息
	}
	return router
}
