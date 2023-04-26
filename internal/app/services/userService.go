package services

import (
	"RedMist/internal/app/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db}
}

func (us UserService) GetUserByUsername(username string) (*models.Users, error) {
	var user models.Users
	err := us.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (us UserService) CreateUser(username, password string) (*models.Users, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.Users{
		Username: username,
		Password: string(passwordHash),
	}
	if err := us.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (us UserService) GetUsers() ([]*models.Users, error) {
	var users []*models.Users
	err := us.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
