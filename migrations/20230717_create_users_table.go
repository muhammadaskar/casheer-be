package main

import (
	"log"

	"github.com/joho/godotenv"
	// "github.com/muhammadaskar/casheer-be/app/product"
	"github.com/muhammadaskar/casheer-be/domains"
	"github.com/muhammadaskar/casheer-be/infrastructures/mysql_driver"
	// "github.com/muhammadaskar/casheer-be/app/product"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db := mysql_driver.InitDatabase()

	db.AutoMigrate(&domains.UserImage{})
	if err != nil {
		panic("Failed to migrate database")
	}
}
