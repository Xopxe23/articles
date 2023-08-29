package service

import "github.com/xopxe23/articles/internal/domain"

type PasswordHasher interface {
	Hash(password string) (string, error)
}

type AuthRepository interface {
	CreateUser(input domain.User) error
}

type AuthService struct {
	authRepos AuthRepository
	hasher    PasswordHasher
}

func NewAuthService(repos AuthRepository, hasher PasswordHasher) *AuthService {
	return &AuthService{authRepos: repos, hasher: hasher}
}

func (s AuthService) SignUp(input domain.User) error {
	password, err := s.hasher.Hash(input.Password)
	if err != nil {
		return err
	}
	return s.authRepos.CreateUser(domain.User{
		Name:     input.Name,
		Surname:  input.Surname,
		Email:    input.Email,
		Password: password,
	})
}
