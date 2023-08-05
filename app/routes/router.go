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

	// CORS MIDDLEWARE
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://38.47.69.131:2000", "http://127.0.0.1:2000"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowHeaders = []string{"Content-Type", "Authorization"}
	config.AllowCredentials = true

	// router.Use(cors.New(config))
	router.Use(cors.New(config))

	// router.Use(CORSMiddleware())

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": true,
			"message": "hello world",
		})
	})

	// api := router.Group("api/v1")
	// {
	// 	// Middleware CORS
	// 	// api.Use(CORSMiddleware())

	// 	api.GET("/", func(c *gin.Context) {
	// 		c.JSON(200, gin.H{
	// 			"success": true,
	// 			"message": "this is api for casheer app",
	// 		})
	// 	})

	// 	api.POST("/auth/register", userHandler.Register)
	// 	api.POST("/auth/login", userHandler.Login)

	// 	// category := api.Group("category")
	// 	// {
	// 	// category.Use(cors.New(config))
	// 	api.GET("/category", authMiddleware, categoryHandler.FindAll)
	// 	api.GET("/category/:id", authMiddleware, categoryHandler.FindById)
	// 	api.POST("/category", authAdminMiddleware, categoryHandler.Create)
	// 	api.PUT("/category/:id", authAdminMiddleware, categoryHandler.Update)
	// 	// }

	// 	// notification := api.Group("notification")
	// 	// {
	// 	// 	// notification.Use(cors.New(config))
	// 	api.GET("/notification", authAdminMiddleware, notificationHandler.FindAll)
	// 	// }
	// }
	api := router.Group("api/v1")

	api.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": true,
			"message": "this is api for casheer app",
		})
	})
	api.POST("/auth/register", userHandler.Register)
	api.POST("/auth/login", userHandler.Login)

	categoryRouter := api.Group("category")

	categoryRouter.GET("/", authMiddleware, categoryHandler.FindAll)
	categoryRouter.GET("/:id", authMiddleware, categoryHandler.FindById)
	categoryRouter.POST("/", authAdminMiddleware, categoryHandler.Create)
	categoryRouter.PUT("/:id", authAdminMiddleware, categoryHandler.Update)

	api.GET("/notification", authAdminMiddleware, notificationHandler.FindAll)

	return router
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Define the list of allowed origins
		// allowedOrigins := []string{
		// 	"http://127.0.0.1",
		// }

		// origin := c.Request.Header.Get("Origin")
		// for _, allowedOrigin := range allowedOrigins {
		// 	if origin == allowedOrigin {
		// 		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		// 		break
		// 	}
		// }
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:2000")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, HEAD")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
