package routes

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	categoryDelivery "github.com/muhammadaskar/casheer-be/app/category/delivery/http"
	categoryRepo "github.com/muhammadaskar/casheer-be/app/category/repository/mysql"
	categoryUseCase "github.com/muhammadaskar/casheer-be/app/category/usecase"
	discountDelivery "github.com/muhammadaskar/casheer-be/app/discount/delivery/http"
	discountRepo "github.com/muhammadaskar/casheer-be/app/discount/repository/mysql"
	discountUseCase "github.com/muhammadaskar/casheer-be/app/discount/usecase"
	memberDelivery "github.com/muhammadaskar/casheer-be/app/member/delivery/http"
	memberRepo "github.com/muhammadaskar/casheer-be/app/member/repository/mysql"
	memberUseCase "github.com/muhammadaskar/casheer-be/app/member/usecase"
	notificationDelivery "github.com/muhammadaskar/casheer-be/app/notification/delivery/http"
	notificationRepo "github.com/muhammadaskar/casheer-be/app/notification/repository/mysql"
	notificationUseCase "github.com/muhammadaskar/casheer-be/app/notification/usecase"
	productDelivery "github.com/muhammadaskar/casheer-be/app/product/delivery/http"
	productRepo "github.com/muhammadaskar/casheer-be/app/product/repository/mysql"
	productUseCase "github.com/muhammadaskar/casheer-be/app/product/usecase"
	transactionDelivery "github.com/muhammadaskar/casheer-be/app/transaction/delivery/http"
	transactionRepo "github.com/muhammadaskar/casheer-be/app/transaction/repository/mysql"
	memberUsecase "github.com/muhammadaskar/casheer-be/app/transaction/usecase"
	userDelivery "github.com/muhammadaskar/casheer-be/app/user/delivery/http"
	userRepo "github.com/muhammadaskar/casheer-be/app/user/repository/mysql"
	userUseCase "github.com/muhammadaskar/casheer-be/app/user/usecase"
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

	userRepository := userRepo.NewRepository(db)
	notificationRepository := notificationRepo.NewRepository(db)
	categoryRepository := categoryRepo.NewRepository(db)
	productRepository := productRepo.NewRepository(db)
	discountRepository := discountRepo.NewRepository(db)
	memberRepository := memberRepo.NewRepository(db)
	transactionRepository := transactionRepo.NewRepository(db)

	authentication := auth.NewJWTAuth()
	userUseCase := userUseCase.NewUseCase(userRepository)
	notificationUseCase := notificationUseCase.NewUseCase(notificationRepository)
	categoryUseCase := categoryUseCase.NewUseCase(categoryRepository)
	productUseCase := productUseCase.NewUseCase(productRepository)
	discountUseCase := discountUseCase.NewUseCase(discountRepository)
	memberUseCase := memberUseCase.NewUseCase(memberRepository)
	transactionUseCase := memberUsecase.NewUseCase(transactionRepository, memberRepository, productRepository, discountRepository)

	userHandler := userDelivery.NewUserHandler(userUseCase, authentication)
	notificationHandler := notificationDelivery.NewNotificationHandler(notificationUseCase)
	categoryHandler := categoryDelivery.NewCategoryHandler(categoryUseCase)
	productHandler := productDelivery.NewProductHandler(productUseCase)
	discountHandler := discountDelivery.NewDiscountHandler(discountUseCase)
	memberHandler := memberDelivery.NewMemberHandler(memberUseCase)
	transactionHandler := transactionDelivery.NewTransactionHandler(transactionUseCase)

	authMiddleware := auth.AuthMiddleware(authentication, userUseCase)
	authAdminMiddleware := auth.AuthAdminMiddleware(authentication, userUseCase)

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
		api.GET("/products", authMiddleware, productHandler.GetAll)
		api.GET("/product", authMiddleware, productHandler.FindAll)
		api.GET("/product/count", authMiddleware, productHandler.CountProducts)
		api.GET("/product/:id", authAdminMiddleware, productHandler.FindById)
		api.POST("/product", authAdminMiddleware, productHandler.CreateProduct)
		api.PUT("/product/:id", authAdminMiddleware, productHandler.UpdateProduct)
		api.DELETE("/product/:id", authAdminMiddleware, productHandler.DeleteProduct)

		api.GET("/discount", authMiddleware, discountHandler.FindByID)
		api.POST("/discount", authAdminMiddleware, discountHandler.Create)
		api.PUT("/discount", authAdminMiddleware, discountHandler.Update)

		api.GET("/member", authMiddleware, memberHandler.FindAll)
		api.POST("/member", authAdminMiddleware, memberHandler.Create)

		api.POST("transaction", authMiddleware, transactionHandler.CreateTransaction)
	}

	return router
}
