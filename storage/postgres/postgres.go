package postgres

import (
	"log"

	model "github.com/webhook-issue-manager/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Inıt() *gorm.DB {
	dsn := "host=localhost user=postgres password=12345 dbname=issue port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&model.Token{})

	return db
}
