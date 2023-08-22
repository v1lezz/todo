package service

import (
	"github.com/v1lezz/todo"
	"github.com/v1lezz/todo/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.UserEntity) (int, error)
}

type TodoList interface {
}

type TodoTask interface {
}

type Service struct {
	Authorization
	TodoList
	TodoTask
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: newAuthorizationService(repos.Authorization),
	}
}
