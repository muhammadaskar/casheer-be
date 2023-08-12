package routes

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/muhammadaskar/casheer-be/app/category"
	"github.com/muhammadaskar/casheer-be/app/handlers"
	"github.com/muhammadaskar/casheer-be/app/middleware"
	"github.com/muhammadaskar/casheer-be/app/notification"
	"github.com/muhammadaskar/casheer-be/app/product"
	"github.com/muhammadaskar/casheer-be/app/user/delivery/http"
	"github.com/muhammadaskar/casheer-be/app/user/repository/mysql"
	"github.com/muhammadaskar/casheer-be/app/user/usecase"
	"github.com/muhammadaskar/casheer-be/infrastructures/auth"
	"github.com/muhammadaskar/casheer-be/infrastructures/mysql_driver"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	db := mysql_driver.InitDatabase()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	userRepository := mysql.NewRepository(db)
	notificationRepository := notification.NewRepository(db)
	categoryRepository := category.NewRepository(db)
	productRepository := product.NewRepository(db)

	authentication := auth.NewJWTAuth()
	userUseCase := usecase.NewUseCase(userRepository)
	notificationService := notification.NewService(notificationRepository)
	categoryService := category.NewService(categoryRepository)
	productService := product.NewService(productRepository)

	userHandler := http.NewUserHandler(userUseCase, authentication)
	notificationHandler := handlers.NewNotificationHandler(notificationService)
	categoryHandler := handlers.NewCategoryHandler(categoryService)
	productHandler := handlers.NewProductHandler(productService)

	authMiddleware := middleware.AuthMiddleware(authentication, userUseCase)
	authAdminMiddleware := middleware.AuthAdminMiddleware(authentication, userUseCase)

	// CORS MIDDLEWARE
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://38.47.69.131:2000", "http://127.0.0.1:2000"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowHeaders = []string{"Content-Type", "Authorization"}
	config.AllowCredentials = true

	router.Use(cors.New(config))

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": true,
			"message": "hello world",
		})
	})

	api := router.Group("api/v1")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"success": true,
				"message": "this is api for casheer app",
			})
		})

		// AUTH
		api.POST("/auth/register", userHandler.Register)
		api.POST("/auth/login", userHandler.Login)

		// CATEGORY
		api.GET("/category", authMiddleware, categoryHandler.FindAll)
		api.GET("/category/:id", authMiddleware, categoryHandler.FindById)
		api.POST("/category", authAdminMiddleware, categoryHandler.Create)
		api.PUT("/category/:id", authAdminMiddleware, categoryHandler.Update)

		// NOTIFICATION
		api.GET("/notification", authAdminMiddleware, notificationHandler.FindAll)

		// PRODUCT
		api.GET("/product", authMiddleware, productHandler.FindAll)
		api.POST("/product", authAdminMiddleware, productHandler.CreateProduct)

	}

	return router
}
