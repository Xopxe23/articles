package service

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/xopxe23/articles/internal/domain"
)

type PasswordHasher interface {
	Hash(password string) (string, error)
}

type AuthRepository interface {
	CreateUser(input domain.User) error
	GetByCredintials(input domain.SignInInput) (domain.User, error)
	AddRefreshToken(token domain.RefreshSession) error
}

type AuthService struct {
	authRepos  AuthRepository
	hasher     PasswordHasher
	hmacSecret []byte
}

func NewAuthService(repos AuthRepository, hasher PasswordHasher, secret []byte) *AuthService {
	return &AuthService{authRepos: repos, hasher: hasher, hmacSecret: secret}
}

func (s *AuthService) SignUp(input domain.User) error {
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

func (s *AuthService) SignIn(input domain.SignInInput) (string, string, error) {
	password, err := s.hasher.Hash(input.Password)
	if err != nil {
		return "", "", err
	}
	input.Password = password
	user, err := s.authRepos.GetByCredintials(input)
	if err != nil {
		return "", "", err
	}
	return s.generateTokens(user.ID)
}

func (s *AuthService) generateTokens(userId int) (string, string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": strconv.Itoa(userId),
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(5 * time.Hour).Unix(),
	})
	accessToken, err := t.SignedString(s.hmacSecret)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := newRefreshToken()
	if err != nil {
		return "", "", err
	}

	if err := s.authRepos.AddRefreshToken(domain.RefreshSession{
		UserId: userId,
		Token: refreshToken,
		ExpiresAt: time.Now().Add(3*time.Hour),
	}); err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}

func newRefreshToken() (string, error) {
	b := make([]byte, 32)
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	if _, err := r.Read(b); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", b), nil
}