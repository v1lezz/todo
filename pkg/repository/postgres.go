package repository

import (
	"fmt"
	"github.com/v1lezz/todo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	DBName   string
	SSLMode  string
}

func InitPostgresDB(cfg DBConfig) (*gorm.DB, error) {
	db, err := gorm.Open(
		postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
			cfg.Host,
			cfg.Username,
			cfg.Password,
			cfg.DBName,
			cfg.Port,
			cfg.SSLMode,
		)),
		&gorm.Config{},
	)
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(
		&todo.UserEntity{},
		&todo.TaskEntity{},
		&todo.CommentEntity{},
		&todo.FileEntity{},
		&todo.TaskEntity{},
		&todo.CategoryEntity{},
	)
	if err != nil {
		return nil, err
	}
	return db, err
}
