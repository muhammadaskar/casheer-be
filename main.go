package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/muhammadaskar/casheer-be/app/routes"
)

type Config struct {
	AllowOrigins []string `json:"ALLOW_ORIGINS"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("SERVER_PORT")
	// clientIp := os.Getenv("CLIENT_IP")

	router := gin.Default()
	config := cors.DefaultConfig()

	var conf Config
	allowOriginsString := os.Getenv("ALLOW_ORIGINS")
	err = json.Unmarshal([]byte(allowOriginsString), &conf.AllowOrigins)
	if err != nil {
		log.Fatal("Error parsing ALLOW_ORIGINS:", err)
		return
	}

	config.AllowOrigins = conf.AllowOrigins
	config.AddAllowHeaders("Access-Control-Allow-Origin")
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowCredentials = true
	config.ExposeHeaders = []string{"Content-Length"}

	router.Use(cors.New(config))
	routes.NewRouter(router)
	router.Run(":" + port)
}
