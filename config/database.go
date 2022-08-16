package config

import (
	"fmt"

	"backend.api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

const (
	HOST     = "0.0.0.0"
	PORT     = "5432"
	USER     = "todo"
	PASSWORD = "password"
	DATABASE = "todo"
	TIMEZONE = "Asia/Bangkok"
)

func ConnectDatabase() {
	dns := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable TimeZone=%s", USER, PASSWORD, HOST, PORT, DATABASE, TIMEZONE)
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		panic("Can't connect to database")
	}

	db.AutoMigrate(&models.Todo{})

	DB = db
}
