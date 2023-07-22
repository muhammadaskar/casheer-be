package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/muhammadaskar/casheer-be/app/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("SERVER_PORT")
	// ORIGIN_PROD := os.Getenv("ALLOW_ORIGIN_PROD")
	// ORIGIN_DEV := os.Getenv("ALLOW_ORIGIN_DEV")

	router := gin.Default()
	config := cors.DefaultConfig()

	config.AllowOrigins = []string{"http://38.47.69.131:3000", "http://127.0.0.1:3000"}
	config.AddAllowHeaders("Access-Control-Allow-Origin")
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowCredentials = true
	config.ExposeHeaders = []string{"Content-Length"}

	router.Use(cors.New(config))
	routes.NewRouter(router)
	router.Run(":" + port)
}
