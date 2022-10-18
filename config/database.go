package config

import (
	"MyGram/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	user     = "postgres"
	password = "8525"
	dbPort   = "5432"
	dbName   = "myGorm"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() *gorm.DB {
	config := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=disable", host, user, password, dbPort, dbName)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database : ", err)
	}
	db.Debug().AutoMigrate(models.Users{})
	return db
}
