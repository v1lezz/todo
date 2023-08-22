package todo

import (
	"gorm.io/gorm"
	"time"
)

type TaskEntity struct {
	gorm.Model
	TaskId      uint
	Deadline    time.Time
	Name        string
	Categories  []*CategoryEntity `gorm:"many2many:task_categories"`
	Comments    []CommentEntity   `gorm:"foreignKey:TaskId"`
	Priority    uint
	IsCompleted bool
}
