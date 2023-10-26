package auth

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func SetupCORS() gin.HandlerFunc {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// allowOriginLocal := os.Getenv("ALLOW_ORIGIN_LOCAL")
	// allowOriginDev := os.Getenv("ALLOW_ORIGIN_DEV")
	// allowOriginProd := os.Getenv("ALLOW_ORIGIN_PROD")

	config := cors.DefaultConfig()
	// config.AllowOrigins = []string{allowOriginDev, allowOriginProd}
	config.AllowOrigins = []string{"http://38.47.69.131:3000", "http://38.47.69.131:2000", "http://127.0.0.1:2000", "http://38.47.69.131"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowHeaders = []string{"Content-Type", "Authorization"}
	config.AllowCredentials = true
	return cors.New(config)
}
