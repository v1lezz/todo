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

func (r *AuthorizationGORM) CreateUser(user todo.UserEntity) (int, error) {
	if err := r.db.Create(&user); err.Error != nil {
		return 0, err.Error
	}
	return int(user.ID), nil
}

func (r *AuthorizationGORM) GetUser(username, password string) (todo.UserEntity, error) {
	user := todo.UserEntity{
		Username: username,
		Password: password,
	}
	if err := r.db.First(&user); err.Error != nil {
		return todo.UserEntity{}, err.Error
	}
	return user, nil
}
