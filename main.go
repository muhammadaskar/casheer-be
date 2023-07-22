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
	clientIp := os.Getenv("CLIENT_IP")

	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{clientIp}
	config.AddAllowHeaders("Access-Control-Allow-Origin")

	router.Use(cors.New(config))

	// router.Use(cors.New(cors.Config{
	// 	// AllowOrigins:     []string{clientIp},
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
	// 	AllowHeaders:     []string{"Origin"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// }))
	routes.NewRouter(router)
	router.Run(":" + port)
}
