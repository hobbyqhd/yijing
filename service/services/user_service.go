package services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/hobbyqhd/yijing/service/config"
	"github.com/hobbyqhd/yijing/service/models"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) Register(username, password, email, nickname string) error {
	// 检查用户名是否已存在
	var existingUser models.User
	result := config.DB.Where("username = ?", username).First(&existingUser)
	if result.RowsAffected > 0 {
		return errors.New("用户名已存在")
	}

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// 创建新用户
	user := models.User{
		Username: username,
		Password: string(hashedPassword),
		Email:    email,
		Nickname: nickname,
	}

	result = config.DB.Create(&user)
	return result.Error
}

func (s *UserService) Login(username, password string) (string, error) {
	// 查找用户
	var user models.User
	result := config.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return "", errors.New("用户不存在")
	}

	// 验证密码
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("密码错误")
	}

	// 生成JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":   user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(config.GetEnv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *UserService) GetUserInfo(userId uint) (*models.User, error) {
	var user models.User
	result := config.DB.First(&user, userId)
	if result.Error != nil {
		return nil, errors.New("用户不存在")
	}

	// 清除敏感信息
	user.Password = ""
	return &user, nil
}

func (s *UserService) UpdateUserInfo(userId uint, nickname, email, avatar string) error {
	result := config.DB.Model(&models.User{}).Where("id = ?", userId).Updates(map[string]interface{}{
		"nickname": nickname,
		"email":    email,
		"avatar":   avatar,
	})
	return result.Error
}
