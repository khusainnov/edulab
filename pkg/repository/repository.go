package repository

import "github.com/khusainnov/edulab/internal/entity/user"

type Authorization interface {
	CreateUser(u user.User) (int, error)
	GetUser(username, password string) (user.User, error)
}

type Repository struct {
	Authorization
}

func NewRepository() *Repository {
	return &Repository{}
}