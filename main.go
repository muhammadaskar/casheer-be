package main

import (
	"github.com/muhammadaskar/casheer-be/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.NewRouter(router)
	router.Run(":3030")
}
