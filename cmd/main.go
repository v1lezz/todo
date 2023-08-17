package main

import (
	"github.com/v1lezz/todo/pkg/config/db"
	comments "github.com/v1lezz/todo/pkg/entities/CommentEntity/models"
	files "github.com/v1lezz/todo/pkg/entities/FileEntity/models"
	"log"
	"time"
)

func main() {
	db, err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	db.Client.Create(&comments.CommentEntity{
		Deadline: time.Now(),
		Payload:  "",
		Files: []files.FileEntity{
			files.FileEntity{LinkToGet: "google.com"},
		},
	})

	//db.Client.Create(&comments.CommentEntity{
	//	gorm.Model{
	//		ID: 0,
	//		CreatedAt: time.Now(),
	//		UpdatedAt: time.Now(),
	//		DeletedAt: gorm.DeletedAt{
	//			Time: nil,
	//			Valid: false,
	//		},
	//	},
	//})
	//router := gin.Default()
	//err := router.Run("localhost:8080")
	//if err != nil {
	//	log.Fatal(err)
	//}
}
