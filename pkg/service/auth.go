package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/v1lezz/todo"
	"github.com/v1lezz/todo/pkg/repository"
)

const salt = "fdasfadsgfahehfdsg156641635asgreasgdsfg"

type AuthorizationService struct {
	repos repository.Authorization
}

func newAuthorizationService(repos repository.Authorization) *AuthorizationService {
	return &AuthorizationService{repos: repos}
}

func (service *AuthorizationService) CreateUser(user todo.UserEntity) (int, error) {
	user.Password = service.hashPassword(user.Password)
	return service.repos.CreateUser(user)
}

func (service *AuthorizationService) hashPassword(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
