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

	port := os.Getenv("SERVER_PORT_DEV")
	ORIGIN_PROD := os.Getenv("ALLOW_ORIGIN_PROD")
	ORIGIN_DEV := os.Getenv("ALLOW_ORIGIN_DEV")

	router := gin.Default()
	config := cors.DefaultConfig()

	config.AllowOrigins = []string{ORIGIN_PROD, ORIGIN_DEV}
	config.AddAllowHeaders("Access-Control-Allow-Origin")
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowCredentials = true
	config.ExposeHeaders = []string{"Content-Length"}

	router.Use(cors.New(config))
	routes.NewRouter(router)
	router.Run(":" + port)
}
