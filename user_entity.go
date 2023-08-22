package todo

import "gorm.io/gorm"

type UserEntity struct {
	gorm.Model
	Name             string       `json:"name" binding:"required"`
	Username         string       `json:"username" binding:"required"`
	Email            *string      `json:"email" binding:"required"`
	IsEmailConfirmed bool         `json:"emailconf"`
	Password         string       `json:"password" binding:"required"`
	Tasks            []TaskEntity `gorm:"foreignKey:TaskId" json:"tasks"`
}
