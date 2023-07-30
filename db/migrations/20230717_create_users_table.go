package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/muhammadaskar/casheer-be/app/notification"

	// "github.com/muhammadaskar/casheer-be/app/product"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Access environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// Membuat tabel "users"
	// err = db.AutoMigrate(&user.User{})
	// err = db.AutoMigrate(&category.Category{})
	// db.AutoMigrate(&product.Product{})
	err = db.AutoMigrate(&notification.Notification{})
	fmt.Println(err)
	if err != nil {
		panic("Failed to migrate database")
	}
}
