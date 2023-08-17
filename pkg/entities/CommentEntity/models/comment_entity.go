package models

import (
	files "github.com/v1lezz/todo/pkg/entities/FileEntity/models"
	"gorm.io/gorm"
	"time"
)

type CommentEntity struct {
	gorm.Model
	TaskId   uint
	Deadline time.Time
	Payload  string
	Files    []files.FileEntity `gorm:"foreignKey:UserId"`
}
