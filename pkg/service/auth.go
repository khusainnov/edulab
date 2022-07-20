package service

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/khusainnov/edulab/internal/entity/user"
	"github.com/khusainnov/edulab/pkg/repository"
)

const (
	salt       = "423r2wedc2r32recdqw"
	tokenTTL   = time.Hour * 12
	signingKey = "4hgb53kh4b5h3g4b345j"
)

type AuthService struct {
	repos repository.Authorization
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func NewAuthService(repos repository.Authorization) *AuthService {
	return &AuthService{repos: repos}
}

func (as *AuthService) CreateUser(u user.User) (int, error) {
	u.Password = generatePasswordHash(u.Password)
	return as.repos.CreateUser(u)
}

func (as *AuthService) GenerateToken(login, password string) (string, error) {
	u, err := as.repos.GetUser(login, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		u.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
