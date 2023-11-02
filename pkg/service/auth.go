package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/v1lezz/todo"
	"github.com/v1lezz/todo/pkg/repository"
	"time"
)

const (
	salt       = "fdasfadsgfahehfdsg156641635asgreasgdsfg"
	signingKey = "fasdgfASDFGASGREAGR3245135HVgafdsgafdg87fdg1"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId uint `json:"user_id"`
}

type AuthorizationService struct {
	repos repository.Authorization
}

func newAuthorizationService(repos repository.Authorization) *AuthorizationService {
	return &AuthorizationService{repos: repos}
}

func (s *AuthorizationService) CreateUser(user todo.UserEntity) (int, error) {
	user.Password = hashPassword(user.Password)
	return s.repos.CreateUser(user)
}

func (s *AuthorizationService) GenerateToken(username, password string) (string, error) {
	user, err := s.repos.GetUser(username, hashPassword(password))
	if err != nil {
		return "", nil
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId: user.ID,
	})

	return token.SignedString([]byte(signingKey))
}

func hashPassword(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
