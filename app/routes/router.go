package routes

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/muhammadaskar/casheer-be/app/auth"
	"github.com/muhammadaskar/casheer-be/app/category"
	"github.com/muhammadaskar/casheer-be/app/config"
	"github.com/muhammadaskar/casheer-be/app/handlers"
	"github.com/muhammadaskar/casheer-be/app/middleware"
	"github.com/muhammadaskar/casheer-be/app/notification"
	"github.com/muhammadaskar/casheer-be/app/user"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	db := config.InitDatabase()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	userRepository := user.NewRepository(db)
	notificationRepository := notification.NewRepository(db)
	categoryRepository := category.NewRepository(db)

	authService := auth.NewService()
	userService := user.NewService(userRepository)
	notificationService := notification.NewService(notificationRepository)
	categoryService := category.NewService(categoryRepository)

	userHandler := handlers.NewUserHandler(userService, authService)
	notificationHandler := handlers.NewNotificationHandler(notificationService)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	authMiddleware := middleware.AuthMiddleware(authService, userService)
	authAdminMiddleware := middleware.AuthAdminMiddleware(authService, userService)

	api := router.Group("api/v1")
	{
		// Middleware CORS
		config := cors.DefaultConfig()
		config.AllowOrigins = []string{"http://38.47.69.131:2000", "http://127.0.0.1:2000"}
		config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
		config.AllowHeaders = []string{"Authorization", "Content-Type"}
		config.AllowCredentials = true
		api.Use(cors.New(config))

		api.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"success": true,
				"message": "hello world",
			})
		})

		api.POST("auth/register", userHandler.Register)
		api.POST("auth/login", userHandler.Login)

		category := api.Group("category")
		{
			category.GET("/", authMiddleware, categoryHandler.FindAll)
			category.GET("/:id", authMiddleware, categoryHandler.FindById)
			category.POST("/", authAdminMiddleware, categoryHandler.Create)
			category.PUT("/:id", authAdminMiddleware, categoryHandler.Update)
		}

		notification := api.Group("notification")
		{
			notification.GET("/", authAdminMiddleware, notificationHandler.FindAll)
		}
	}

	return router
}
