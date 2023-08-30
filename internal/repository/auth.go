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

func (r *AuthRepository) GetByCredintials(input domain.SignInInput) (domain.User, error) {
	var user domain.User
	query := "SELECT * FROM users WHERE email=$1 and password=$2"
	if err := r.DB.Get(&user, query, input.Email, input.Password); err != nil {
		return user, err
	}
	return user, nil
}

func (r *AuthRepository) AddRefreshToken(token domain.RefreshSession) error {
	_, err := r.DB.Exec("INSERT INTO refresh_tokens(user_id, token, expires_at) VALUES ($1, $2, $3)",
		token.UserId, token.Token, token.ExpiresAt)
	return err
}

func (r *AuthRepository) GetToken(token string) (domain.RefreshSession, error) {
	var session domain.RefreshSession
	if err := r.DB.Get(&session, "SELECT * FROM refresh_tokens WHERE token=$1", token); err != nil {
		return session, err
	}
	_, err := r.DB.Exec("DELETE FROM refresh_tokens WHERE user_id=$1", session.UserId)
	return session, err
}
