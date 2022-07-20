package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/khusainnov/edulab/internal/entity/user"
)

type Authorization interface {
	CreateUser(u user.User) (int, error)
	GetUser(login, password string) (user.User, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
