package services

import (
	"Z02zz/internal/models"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// UserService 定义用户相关的业务逻辑
type UserService struct {
	DB *gorm.DB
}

// NewUserService 创建一个新的 UserService 实例
func NewUserService(db *gorm.DB) *UserService {
	return &UserService{DB: db}
}

// RegisterUser 注册新用户，并将密码进行哈希处理
func (service *UserService) RegisterUser(username, email, password string) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
	}

	if err := service.DB.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// AuthenticateUser 验证用户名和密码是否匹配
func (service *UserService) AuthenticateUser(username, password string) (*models.User, error) {
	var user models.User
	if err := service.DB.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户不存在")
		}
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("密码错误")
	}

	return &user, nil
}

// UpdateUser 更新用户信息
func (service *UserService) UpdateUser(user *models.User) error {
	return service.DB.Save(user).Error
}
