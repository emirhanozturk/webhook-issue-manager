package postgres

import (
	"fmt"
	"log"

	"github.com/webhook-issue-manager/config"
	model "github.com/webhook-issue-manager/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Inıt() *gorm.DB {
	config := config.Config("./config.yaml")
	host := config.Hostname
	port := config.Port
	database := config.Database
	user := config.User
	password := config.Password
	dsn := fmt.Sprintf("host=%s user=%s password=%d dbname=%s port=%d sslmode=disable", host, user, password, database, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&model.Token{}, &model.Assignee{}, &model.Issue{}, model.Comment{}, &model.Attachment{})

	return db
}
