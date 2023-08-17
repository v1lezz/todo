package models

import (
	categories "github.com/v1lezz/todo/pkg/entities/CategoryEntity/models"
	comment "github.com/v1lezz/todo/pkg/entities/CommentEntity/models"
	"gorm.io/gorm"
	"time"
)

type TaskEntity struct {
	gorm.Model
	TaskId      uint
	Deadline    time.Time
	Name        string
	Categories  []*categories.CategoryEntity `gorm:"many2many:task_categories"`
	Comments    []comment.CommentEntity      `gorm:"foreignKey:TaskId"`
	Priority    uint
	IsCompleted bool
}
