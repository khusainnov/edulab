package service

import (
	"github.com/khusainnov/edulab/internal/entity/user"
	"github.com/khusainnov/edulab/pkg/repository"
)

type Authorization interface {
	CreateUser(u user.User) (int, error)
	GenerateToken(username, password string) (string, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos),
	}
}