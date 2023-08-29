package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/xopxe23/articles/internal/domain"
)

type AuthRepository struct {
	DB *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{DB: db}
}

func (r *AuthRepository) CreateUser(input domain.User) error {
	_, err := r.DB.Exec("INSERT INTO users(name, surname, email, password) VALUES ($1, $2, $3, $4)",
		input.Name, input.Surname, input.Email, input.Password)
	return err
}
