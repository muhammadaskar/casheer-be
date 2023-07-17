package routes

import (
	"github.com/gin-gonic/gin"
)

func NewRouter(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": true,
			"message": "pong",
		})
	})

	api := router.Group("api/v1")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"success": true,
				"message": "hello world",
			})
		})
	}
}
