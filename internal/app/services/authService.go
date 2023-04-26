package services

import (
	"errors"
	"github.com/go-chi/jwtauth/v5"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userService *UserService
	tokenAuth   *jwtauth.JWTAuth
}

type TokenClaims struct {
	jwtauth.JWTAuth
	UserID uint `json:"user_id"`
}

func NewAuthService(userService *UserService, jwtSecret []byte) *AuthService {
	return &AuthService{
		userService: userService,
		tokenAuth:   jwtauth.New("HS256", jwtSecret, nil),
	}
}

func (a *AuthService) Register(name, email, password string) (*User, error) {
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return nil, err
	}
	user, err := a.userService.CreateUser(name, email, hashedPassword)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (a *AuthService) Login(email, password string) (string, error) {
	user, err := a.userService.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("Invalid email or password")
	}
	if !CheckPasswordHash(password, user.Password) {
		return "", errors.New("Invalid email or password")
	}
	token, err := a.createToken(user.ID, time.Now().Add(time.Hour*24))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (a *AuthService) createToken(userID uint, expires time.Time) (string, error) {
	tokenAuth := a.tokenAuth.New("access_token", expires.Unix())
	tokenAuth.SetField("user_id", userID)
	tokenString, err := tokenAuth.Encode()
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
