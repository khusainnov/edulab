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
	salt       = "45njh5n34jh5jh3n4j2"
	signingKey = "4hbj4b5k34b5hg3bh42b4h234h2bh"
	tokenTTL   = time.Hour * 12
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repos *repository.Repository
}

func NewAuthService(repos *repository.Repository) *AuthService {
	return &AuthService{repos: repos}
}

func (as *AuthService) CreateUser(u user.User) (int, error) {
	u.Password = generatePasswordHash(u.Password)

	return as.repos.CreateUser(u)
}

func (as *AuthService) GenerateToken(username, password string) (string, error) {
	u, err := as.repos.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims{
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
