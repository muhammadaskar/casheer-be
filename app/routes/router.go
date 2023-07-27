package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammadaskar/casheer-be/app/auth"
	"github.com/muhammadaskar/casheer-be/app/category"
	"github.com/muhammadaskar/casheer-be/app/config"
	"github.com/muhammadaskar/casheer-be/app/handlers"
	"github.com/muhammadaskar/casheer-be/app/middleware"
	"github.com/muhammadaskar/casheer-be/app/user"
)

func NewRouter(router *gin.Engine) {
	db := config.InitDatabase()

	userRepository := user.NewRepository(db)
	categoryRepository := category.NewRepository(db)

	userService := user.NewService(userRepository)
	authService := auth.NewService()
	categoryService := category.NewService(categoryRepository)

	userHandler := handlers.NewUserHandler(userService, authService)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	authMiddleware := middleware.AuthMiddleware(authService, userService)
	authAdminMiddleware := middleware.AuthAdminMiddleware(authService, userService)

	api := router.Group("api/v1")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"success": true,
				"message": "hello world",
			})
		})

		api.POST("auth/register", userHandler.Register)
		api.POST("auth/login", userHandler.Login)

		// category := api.Group("category")
		// {
		api.GET("/category", authMiddleware, categoryHandler.FindAll)
		api.GET("/category/:id", authMiddleware, categoryHandler.FindById)
		api.POST("/category", authAdminMiddleware, categoryHandler.Create)
		api.PUT("/category/:id", authAdminMiddleware, categoryHandler.Update)
		// }
	}
}
