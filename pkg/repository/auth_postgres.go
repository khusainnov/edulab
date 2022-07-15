package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/khusainnov/edulab/internal/entity/user"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (ap *AuthPostgres) CreateUser(u user.User) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (name, surname, username, email, password_hash) VALUES ($1, $2, $3, $4, $5) RETURNING id;", userTable)
	row := ap.db.QueryRow(query, u.Name, u.Surname, u.Username, u.Email, u.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (ap *AuthPostgres) GetUser(username, password string) (user.User, error) {
	var u user.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2;", userTable)
	err := ap.db.Get(&u, query, username, password)

	return u, err
}