package db

import (
	"fmt"
	"log"
	"myblog/internal/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	var err error
	//配置DSN字符串
	dsn := "root:root@tcp(127.0.0.1:3306)/myblog?charset=utf8mb4&parseTime=True&loc=Local"

	//使用GORM和DSN字符串连接到MYSQL
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	log.Println("Database connected successfully.")

	fmt.Println(*DB)

	// 自动迁移
	err = DB.AutoMigrate(&models.Post{})
	if err != nil {
		log.Fatal("failed to migrate database:", err)
	}

	log.Println("Database connected and migrated successfully.")
}
