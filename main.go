package main

import (
	"gin-introduction/controllers"
	"gin-introduction/infra"
	"gin-introduction/repositories"

	// "gin-introduction/models"

	"gin-introduction/services"

	"github.com/gin-gonic/gin"
)

func main() {
	infra.Initializer()
	db := infra.SetupDB()


	// items := []models.Item{
	// 	{ID: 1, Name: "商品1", Price: 1000, Description: "説明1", SoldOut: false},
	// 	{ID: 2, Name: "商品2", Price: 2000, Description: "説明2", SoldOut: true},
	// 	{ID: 3, Name: "商品3", Price: 3000, Description: "説明3", SoldOut: false},
	// }

	// itemRepository := repositories.NewItemMemoryRepository(items)
	itemRepository := repositories.NewItemRepository(db)
	itemService := services.NewItemService(itemRepository)
	itemController := controllers.NewItemController(itemService)
	
	authRepository := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepository)
	authController := controllers.NewAuthController(authService)

	router := gin.Default()
	itemRouter := router.Group("/items")
	authRouter := router.Group("/auth")
	

	itemRouter.GET("/", itemController.FindAll)
	itemRouter.GET("/:id", itemController.FindById)
	itemRouter.POST("/", itemController.Create)
	itemRouter.PUT("/:id", itemController.Update)
	itemRouter.DELETE("/:id", itemController.Delete)

	authRouter.POST("/signup", authController.SignUp)
	authRouter.POST("/login", authController.Login)
	router.Run("localhost:8080")
}
