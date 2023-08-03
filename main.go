package main

import (
	"log"
	// "net/http"
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

	router := gin.Default()

	// // Middleware CORS
	// config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"http://127.0.0.1:2000"}
	// config.AllowMethods = []string{"GET", "HEAD", "PUT", "PATCH", "POST", "DELETE"}
	// config.AllowHeaders = []string{"Origin", "X-Requested-With", "Content-Type", "Accept", "Authorization"}
	// config.AllowCredentials = true

	// router.Use(cors.New(config))
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://38.47.69.131:2000"},
		AllowMethods:     []string{"GET", "HEAD", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type", "Accept", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	routes.NewRouter(router)
	router.Run(":" + port)
	// http.ListenAndServe(":"+port, router)
}
