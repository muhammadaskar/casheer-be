package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/muhammadaskar/casheer-be/app/routes"
)

func main() {

	router := gin.Default()

	router.Use(cors.Default())
	routes.NewRouter(router)
	router.Run(":3030")
}
