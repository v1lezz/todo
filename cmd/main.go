package main

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/v1lezz/todo"
	"github.com/v1lezz/todo/pkg/handler"
	"github.com/v1lezz/todo/pkg/handler/email"
	"github.com/v1lezz/todo/pkg/repository"
	"github.com/v1lezz/todo/pkg/service"
	"os"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initCFG(); err != nil {
		logrus.Fatalf("error init cfg: %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error init enviroment: %s", err.Error())
	}

	db, err := repository.InitPostgresDB(repository.DBConfig{
		viper.GetString("db.username"),
		os.Getenv("DB_PASSWORD"),
		viper.GetString("db.host"),
		viper.GetString("db.port"),
		viper.GetString("db.dbname"),
		viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("error start db: %s", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handler := handler.NewHandler(services, email.NewEmailSmtp(email.NewConfig(
		viper.GetString("email.mail"),
		os.Getenv("EMAIL_PASSWORD"),
		viper.GetString("email.host"),
		viper.GetString("email.port"),
	)))
	server := new(todo.Server)
	if err = server.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		logrus.Fatalf("error start server: %s", err.Error())
	}
}

func initCFG() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
