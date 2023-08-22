package repository

import (
	"github.com/v1lezz/todo"
	"gorm.io/gorm"
)

type AuthorizationGORM struct {
	db *gorm.DB
}

func NewAuthorizationGORM(db *gorm.DB) *AuthorizationGORM {
	return &AuthorizationGORM{db: db}
}

func (repository *AuthorizationGORM) CreateUser(user todo.UserEntity) (int, error) {
	if err := repository.db.Create(&user); err.Error != nil {
		return 0, err.Error
	}
	return int(user.ID), nil
}
