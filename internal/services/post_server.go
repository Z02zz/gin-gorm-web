package services

import (
	"Z02zz/internal/models"
	"errors"

	"gorm.io/gorm"
)

// PostService defines the business logic for managing blog posts
type PostService struct {
	DB *gorm.DB
}

// NewPostService creates a new instance of PostService
func NewPostService(db *gorm.DB) *PostService {
	return &PostService{DB: db}
}

// GetAllPosts retrieves all blog posts
func (service *PostService) GetAllPosts() ([]models.Post, error) {
	var posts []models.Post
	if err := service.DB.Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

// GetPostByID retrieves a post by its ID
func (service *PostService) GetPostByID(id uint) (*models.Post, error) {
	var post models.Post
	if err := service.DB.First(&post, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // 如果没有找到，返回nil而不是错误
		}
		return nil, err
	}
	return &post, nil
}

// CreatePost creates a new blog post
func (service *PostService) CreatePost(post *models.Post) error {
	if post.Title == "" || post.Content == "" {
		return errors.New("title and content cannot be empty")
	}
	return service.DB.Create(post).Error
}

// UpdatePost updates an existing post
func (service *PostService) UpdatePost(post *models.Post) error {
	existingPost, err := service.GetPostByID(post.ID)
	if err != nil {
		return err
	}
	if existingPost == nil {
		return errors.New("post not found")
	}
	return service.DB.Save(post).Error
}

// DeletePost deletes a post by its ID (soft delete)
func (service *PostService) DeletePost(id uint) error {
	return service.DB.Delete(&models.Post{}, id).Error
}
