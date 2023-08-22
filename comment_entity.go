package todo

import (
	"gorm.io/gorm"
	"time"
)

type CommentEntity struct {
	gorm.Model
	TaskId   uint
	Deadline time.Time
	Payload  string
	Files    []FileEntity `gorm:"foreignKey:UserId"`
}
