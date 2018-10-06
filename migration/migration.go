package migration

import (
	"github.com/jorgeAM/comments/config"
	"github.com/jorgeAM/comments/models"
)

// Migrate -> migraciones
func Migrate() {
	db := config.GetConnection()
	defer db.Close()
	db.CreateTable(&models.User{})
	db.CreateTable(&models.Comment{})
	db.CreateTable(&models.Vote{})
	//para que solo en un registro puede haber un voto por persona
	db.Model(&models.Vote{}).AddUniqueIndex("comment_id_user_id_unique")
}
