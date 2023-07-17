package main

import (
	"github.com/muhammadaskar/casheer-be/app/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:default@tcp(localhost:3306)/db_dev_casheer?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// Membuat tabel "users"
	err = db.AutoMigrate(&user.User{})
	if err != nil {
		panic("Failed to migrate database")
	}
}
