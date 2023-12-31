package routes

import (
	"github.com/gin-gonic/gin"
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
	storeDelivery "github.com/muhammadaskar/casheer-be/app/store/delivery/http"
	storeRepo "github.com/muhammadaskar/casheer-be/app/store/repository/mysql"
	storeUseCase "github.com/muhammadaskar/casheer-be/app/store/usecase"
	transactionDelivery "github.com/muhammadaskar/casheer-be/app/transaction/delivery/http"
	transactionRepo "github.com/muhammadaskar/casheer-be/app/transaction/repository/mysql"
	memberUsecase "github.com/muhammadaskar/casheer-be/app/transaction/usecase"
	userDelivery "github.com/muhammadaskar/casheer-be/app/user/delivery/http"
	userRepo "github.com/muhammadaskar/casheer-be/app/user/repository/mysql"
	userUseCase "github.com/muhammadaskar/casheer-be/app/user/usecase"
	userImageDelivery "github.com/muhammadaskar/casheer-be/app/user_image/delivery/http"
	userImageRepo "github.com/muhammadaskar/casheer-be/app/user_image/repository/mysql"
	userImageUseCase "github.com/muhammadaskar/casheer-be/app/user_image/usecase"
	"github.com/muhammadaskar/casheer-be/infrastructures/auth"
	"github.com/muhammadaskar/casheer-be/infrastructures/mysql_driver"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	db := mysql_driver.InitDatabase()

	userRepository := userRepo.NewRepository(db)
	notificationRepository := notificationRepo.NewRepository(db)
	categoryRepository := categoryRepo.NewRepository(db)
	productRepository := productRepo.NewRepository(db)
	discountRepository := discountRepo.NewRepository(db)
	memberRepository := memberRepo.NewRepository(db)
	transactionRepository := transactionRepo.NewRepository(db)
	storeRepository := storeRepo.NewRepository(db)
	userImageRepository := userImageRepo.NewRepository(db)

	authentication := auth.NewJWTAuth()
	userUseCase := userUseCase.NewUseCase(userRepository, notificationRepository)
	notificationUseCase := notificationUseCase.NewUseCase(notificationRepository)
	categoryUseCase := categoryUseCase.NewUseCase(categoryRepository)
	productUseCase := productUseCase.NewUseCase(productRepository, notificationRepository)
	discountUseCase := discountUseCase.NewUseCase(discountRepository)
	memberUseCase := memberUseCase.NewUseCase(memberRepository)
	transactionUseCase := memberUsecase.NewUseCase(transactionRepository, memberRepository, productRepository, discountRepository)
	storeUseCase := storeUseCase.NewUseCase(storeRepository)
	userImageUseCase := userImageUseCase.NewUseCase(userImageRepository)

	userHandler := userDelivery.NewUserHandler(userUseCase, authentication)
	notificationHandler := notificationDelivery.NewNotificationHandler(notificationUseCase)
	categoryHandler := categoryDelivery.NewCategoryHandler(categoryUseCase)
	productHandler := productDelivery.NewProductHandler(productUseCase)
	discountHandler := discountDelivery.NewDiscountHandler(discountUseCase)
	memberHandler := memberDelivery.NewMemberHandler(memberUseCase)
	transactionHandler := transactionDelivery.NewTransactionHandler(transactionUseCase)
	storeHandler := storeDelivery.NewStoreHandler(storeUseCase)
	userImageHandler := userImageDelivery.NewUserImageHandler(userImageUseCase)

	authMiddleware := auth.AuthMiddleware(authentication, userUseCase)
	authAdminMiddleware := auth.AuthAdminMiddleware(authentication, userUseCase)

	gin.SetMode(gin.ReleaseMode)

	// CORS MIDDLEWARE
	router.Use(auth.SetupCORS())

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": true,
			"message": "hello world!!!",
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
		api.PUT("/notification/:id", authAdminMiddleware, notificationHandler.Update)

		// PRODUCT
		api.GET("/products", authMiddleware, productHandler.GetAll)
		api.GET("/product", authMiddleware, productHandler.FindAll)
		api.GET("/product/deleted", authMiddleware, productHandler.FindAllIsDeleted)
		api.GET("/product/count", authMiddleware, productHandler.CountProducts)
		api.GET("/product/:id", authAdminMiddleware, productHandler.FindById)
		api.POST("/product", authAdminMiddleware, productHandler.CreateProduct)
		api.PUT("/product/:id", authAdminMiddleware, productHandler.UpdateProduct)
		api.PUT("/product/quantity/:id", authAdminMiddleware, productHandler.UpdateProductQuantity)
		api.DELETE("/product/:id", authAdminMiddleware, productHandler.DeleteProduct)

		api.GET("/discount", authMiddleware, discountHandler.FindByID)
		api.POST("/discount", authAdminMiddleware, discountHandler.Create)
		api.PUT("/discount", authAdminMiddleware, discountHandler.Update)

		api.GET("/member", authMiddleware, memberHandler.FindAll)
		api.POST("/member", authAdminMiddleware, memberHandler.Create)
		api.PUT("/member/:id", authAdminMiddleware, memberHandler.Update)

		api.GET("/transaction", authMiddleware, transactionHandler.FindAll)
		api.GET("/transaction/member", authMiddleware, transactionHandler.FindAllMember)
		api.GET("/transaction/:id", authMiddleware, transactionHandler.FindById)
		api.GET("/transaction/member/:id", authMiddleware, transactionHandler.FindMemberById)
		api.GET("/transaction/amount", authMiddleware, transactionHandler.GetAmountOneMonthAgo)
		api.GET("/transaction/item-out", authMiddleware, transactionHandler.GetItemOutOneMonthAgo)
		api.GET("/transaction/count/this-year", authMiddleware, transactionHandler.GetCountTransactionThisYear)
		api.GET("/transaction/amount/this-year", authMiddleware, transactionHandler.GetAmountTransactionThisYear)
		api.POST("/transaction", authMiddleware, transactionHandler.CreateTransaction)

		api.GET("/users", authAdminMiddleware, userHandler.GetUserCasheers)
		api.GET("/users/admin", authAdminMiddleware, userHandler.GetUserAdmin)
		api.PUT("/users/change-to-admin/:id", authAdminMiddleware, userHandler.ChangeToAdmin)
		api.GET("/users/unprocess", authAdminMiddleware, userHandler.GetUsersUnprocess)
		api.GET("/users/rejected", authAdminMiddleware, userHandler.GetUsersRejected)
		api.PUT("/user/activate/:id", authAdminMiddleware, userHandler.Activate)
		api.PUT("/user/reject/:id", authAdminMiddleware, userHandler.Reject)
		api.GET("/user/total-casheer", authAdminMiddleware, userHandler.GetTotalCasheer)
		api.PUT("/user/update/name-or-email", authMiddleware, userHandler.UpdateNameOrEmail)
		api.PUT("/user/update/password", authMiddleware, userHandler.UpdatePassword)

		api.GET("/user/profile-image", authMiddleware, userImageHandler.FindByUserId)
		api.POST("/user/profile-image", authMiddleware, userImageHandler.CreateOrUpdate)

		api.GET("/store", storeHandler.FindOne)
		api.POST("/store", authAdminMiddleware, storeHandler.Create)
		api.PUT("/store", authAdminMiddleware, storeHandler.Update)

	}

	return router
}
