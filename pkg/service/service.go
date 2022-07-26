package service

import (
	"github.com/khusainnov/edulab/internal/entity/user"
	"github.com/khusainnov/edulab/pkg/repository"
)

type Authorization interface {
	CreateUser(u user.User) (int, error)
	GenerateToken(login, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
