package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	PORT     string `mapstructure:"PORT"`
	APP_NAME string `mapstructure:"APP_NAME"`
	DBUser   string `mapstructure:"DB_USER"`
	DBPass   string `mapstructure:"DB_PASSWORD"`
	DBHost   string `mapstructure:"DB_HOST"`
	DBPort   string `mapstructure:"DB_PORT"`
	DBName   string `mapstructure:"DB_NAME"`
}

var AppConfig *Config

func LoadConfig() {
	viper.SetConfigName("config")          // 配置文件名 (不带扩展名)
	viper.SetConfigType("yaml")            // 配置文件类型
	viper.AddConfigPath("internal/config") // 配置文件所在路径
	viper.AutomaticEnv()                   // 允许从环境变量中读取配置

	// 默认值设置
	viper.SetDefault("PORT", "8080")
	viper.SetDefault("APP_NAME", "Z02zz")

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file: %v", err)
	}

	// 将配置数据解析到 Config 结构体
	AppConfig = &Config{}
	if err := viper.Unmarshal(AppConfig); err != nil {
		log.Fatalf("Unable to decode into struct: %v", err)
	}
}
