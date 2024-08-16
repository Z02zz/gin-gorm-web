package models

import (
	"time"

	"gorm.io/gorm"
)

// Post represents a blog post in the database
type Post struct {
	ID        uint           `gorm:"primaryKey"`                 // 主键
	Title     string         `gorm:"type:varchar(255);not null"` // 文章标题
	Content   string         `gorm:"type:text;not null"`         // 文章内容
	AuthorID  uint           `gorm:"not null"`                   // 作者ID，外键关联用户表
	CreatedAt time.Time      `gorm:"autoCreateTime"`             // 创建时间
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`             // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index"`                      // 软删除字段，使用索引
}

// TableName specifies the table name for Post
func (Post) TableName() string {
	return "posts"
}
