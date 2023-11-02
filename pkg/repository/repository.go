package repository

import (
	"github.com/v1lezz/todo"
	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(user todo.UserEntity) (int, error)
	GetUser(username, password string) (todo.UserEntity, error)
}

type TodoList interface {
}

type TodoTask interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoTask
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthorizationGORM(db),
	}
}
