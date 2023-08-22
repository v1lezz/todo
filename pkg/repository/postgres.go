package db

import (
	"github.com/v1lezz/todo"
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
		&todo.UserEntity{},
		&todo.TaskEntity{},
		&todo.CommentEntity{},
		&todo.FileEntity{},
		&todo.TaskEntity{},
		&todo.CategoryEntity{},
	)

}

func InitDB() (*DataBase, error) {
	db := &DataBase{}
	if err := db.ConnectDB(); err != nil {
		return nil, err
	}

	return db, nil
}
