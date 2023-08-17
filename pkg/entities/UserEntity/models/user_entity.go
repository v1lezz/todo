package models

import (
	tasks "github.com/v1lezz/todo/pkg/entities/TaskEntity/models"
	"gorm.io/gorm"
)

type UserEntity struct {
	gorm.Model
	Tasks            []tasks.TaskEntity `gorm:"foreignKey:TaskId"`
	Role             string
	Email            string
	IsEmailConfirmed bool
}
