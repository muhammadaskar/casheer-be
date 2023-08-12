package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/muhammadaskar/casheer-be/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("SERVER_PORT_DEV")
	router := routes.NewRouter()
	router.Run(":" + port)
}
