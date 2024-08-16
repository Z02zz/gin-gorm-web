package database

import (
	"Z02zz/internal/config"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() {
	var err error

	// 从配置中获取数据库连接信息
	dbUser := config.AppConfig.DBUser
	dbPassword := config.AppConfig.DBPass
	dbHost := config.AppConfig.DBHost
	dbPort := config.AppConfig.DBPort
	dbName := config.AppConfig.DBName

	// 检查配置是否为空
	if dbHost == "" || dbPort == "" || dbUser == "" || dbName == "" {
		log.Fatalf("数据库配置无效，请检查配置文件")
	}

	// 构建 DSN
	DSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)

	// 输出 DSN 进行验证
	log.Printf("DSN: %s", DSN)

	// 设置数据库连接
	DB, err = gorm.Open(mysql.Open(DSN), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// 设置连接池
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Failed to get database connection: %v", err)
	}

	sqlDB.SetMaxIdleConns(10)           // 最大空闲连接数
	sqlDB.SetMaxOpenConns(100)          // 最大连接数
	sqlDB.SetConnMaxLifetime(time.Hour) // 连接的最大生命周期

	log.Println("Database connection initialized successfully")
}

func CloseDB() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Failed to get database connection: %v", err)
	}
	sqlDB.Close()
}
