package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammadaskar/casheer-be/app/auth"
	"github.com/muhammadaskar/casheer-be/app/config"
	"github.com/muhammadaskar/casheer-be/app/handlers"
	"github.com/muhammadaskar/casheer-be/app/user"
)

func NewRouter(router *gin.Engine) {
	db := config.InitDatabase()

	userRepository := user.NewRepository(db)

	userService := user.NewService(userRepository)
	authService := auth.NewService()

	userHandler := handlers.NewUserHandler(userService, authService)

	api := router.Group("api/v1")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"success": true,
				"message": "hello world",
			})
		})

		api.POST("auth/register", userHandler.Register)
	}
}
