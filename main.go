package main

import (
	"Z02zz/internal/config"
	"Z02zz/internal/controllers"
	"Z02zz/internal/models"
	"Z02zz/internal/routes"
	"Z02zz/internal/services"
	"Z02zz/pkg/database"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// 加载配置
	config.LoadConfig()

	// 初始化数据库连接
	database.InitDB()
	db := database.DB
	// *gorm.DB 实例
	db.AutoMigrate(&models.Post{})

	// 初始化服务
	postService := services.NewPostService(db)
	userService := services.NewUserService(db)

	// 初始化控制器
	postController := controllers.NewPostController(postService)
	userController := controllers.NewUserController(userService)

	// 设置路由并传递控制器
	router := routes.SetupRouter(postController, userController)

	// 创建服务器
	server := &http.Server{
		Addr:    ":" + config.AppConfig.PORT,
		Handler: router,
	}

	// 在协程中启动服务器
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("服务器启动失败: %v", err)
		}
	}()

	// 捕获系统信号，优雅关闭
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("正在关闭服务器...")

	// 创建5秒超时的上下文
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 优雅关闭服务器（等待正在处理的请求结束）
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("服务器强制关闭: %v", err)
	}

	log.Println("服务器已关闭")

	// 关闭数据库连接
	database.CloseDB()

	log.Println("数据库连接已关闭")
}
