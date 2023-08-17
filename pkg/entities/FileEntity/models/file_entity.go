package models

import "gorm.io/gorm"

type FileEntity struct {
	gorm.Model
	UserId    uint
	LinkToGet string
}
