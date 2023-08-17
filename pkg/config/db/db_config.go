package db

import (
	categories "github.com/v1lezz/todo/pkg/entities/CategoryEntity/models"
	comments "github.com/v1lezz/todo/pkg/entities/CommentEntity/models"
	files "github.com/v1lezz/todo/pkg/entities/FileEntity/models"
	tasks "github.com/v1lezz/todo/pkg/entities/TaskEntity/models"
	users "github.com/v1lezz/todo/pkg/entities/UserEntity/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DataBase struct {
	Client *gorm.DB
}

func (db *DataBase) ConnectDB() (err error) {
	dsn := "host=localhost user=postgres password=postgres dbname=todo port=5432 sslmode=disable"
	db.Client, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	return db.Client.AutoMigrate(
		&users.UserEntity{},
		&tasks.TaskEntity{},
		&comments.CommentEntity{},
		&files.FileEntity{},
		//&tc.TaskCategoryEntity{},
		&tasks.TaskEntity{},
		&categories.CategoryEntity{},
	)

	//return err
}

func InitDB() (*DataBase, error) {
	db := &DataBase{}
	if err := db.ConnectDB(); err != nil {
		return nil, err
	}

	return db, nil
}
