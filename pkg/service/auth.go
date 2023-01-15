package service

import (
	"crypto/sha1"
	"fmt"
	"time"

	todo "github.com/POMBNK/restAPI"
	"github.com/POMBNK/restAPI/pkg/repository"
	"github.com/dgrijalva/jwt-go"
)

// TODO: move to .env file!
const (
	salt       = "sagk@$#@olaed3423hgjs25dklt$^%^ghdaj"
	sign       = "djshf#$giklahl$%kg1234y#$@kdflhgbdsaf9080"
	expireTime = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username string, password string) (string, error) {
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expireTime).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id})

	return token.SignedString([]byte(sign))

}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
