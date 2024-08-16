package models

import (
	"time"

	"gorm.io/gorm"
)

// User represents a user in the database
type User struct {
	ID        uint           `gorm:"primaryKey"`               // 主键
	Username  string         `gorm:"type:varchar(100);unique"` // 用户名，唯一
	Email     string         `gorm:"type:varchar(100);unique"` // 邮箱，唯一
	Password  string         `gorm:"type:varchar(255)"`        // 密码，哈希存储
	CreatedAt time.Time      `gorm:"autoCreateTime"`           // 创建时间
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`           // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index"`                    // 软删除字段
}

// TableName specifies the table name for User
func (User) TableName() string {
	return "users"
}
