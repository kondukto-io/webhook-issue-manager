package postgres

import (
	"fmt"
	"log"

	"github.com/webhook-issue-manager/config"
	model "github.com/webhook-issue-manager/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	appConfig := config.Config("config.yaml")

	var host = appConfig.Host
	var port = appConfig.Port
	var database = appConfig.Database
	var user = appConfig.User
	var password = appConfig.Password
	var dsn = fmt.Sprintf("host=%s user=%s password=%d dbname=%s port=%d sslmode=disable", host, user, password, database, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("db name: %s", db.Name())

	if err := db.AutoMigrate(&model.Token{}, &model.Assignee{}, &model.Issue{}, model.Comment{}, &model.Attachment{}); err != nil {
		log.Fatal(err)
	}

	return db
}
