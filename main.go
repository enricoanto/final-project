package main

import (
	"github.com/enricoanto/final-project/config"
	"github.com/enricoanto/final-project/routes"
	"github.com/enricoanto/final-project/routes/middleware"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"

	// Controller
	categoryController "github.com/enricoanto/final-project/controller/category"
	productController "github.com/enricoanto/final-project/controller/product"
	transactionHistoryController "github.com/enricoanto/final-project/controller/transaction_history"
	userController "github.com/enricoanto/final-project/controller/user"

	// Repository
	categoryRepository "github.com/enricoanto/final-project/repository/category"
	productRepository "github.com/enricoanto/final-project/repository/product"
	transactionHistoryRepository "github.com/enricoanto/final-project/repository/transaction_history"
	userRepository "github.com/enricoanto/final-project/repository/user"

	// Service
	categoryService "github.com/enricoanto/final-project/service/category"
	productService "github.com/enricoanto/final-project/service/product"
	transactionHistoryService "github.com/enricoanto/final-project/service/transaction_history"
	userService "github.com/enricoanto/final-project/service/user"
)

func main() {
	r := gin.Default()
	DB := config.DB

	categoryRepository := categoryRepository.NewRepository(DB)
	productRepository := productRepository.NewRepository(DB)
	transactionHistoryRepository := transactionHistoryRepository.NewRepository(DB)
	userRepository := userRepository.NewRepository(DB)

	userService := userService.NewService(userRepository)
	categoryService := categoryService.NewService(categoryRepository)
	productService := productService.NewService(productRepository)
	transactionHistoryService := transactionHistoryService.NewService(
		transactionHistoryRepository,
		categoryService,
		productService,
		userService,
	)

	userController := userController.NewController(userService)
	categoryController := categoryController.NewController(categoryService)
	productController := productController.NewController(productService)
	transactionHistoryController := transactionHistoryController.NewController(transactionHistoryService)
	middlewareController := middleware.NewController()

	routes.API(
		r,
		userController,
		categoryController,
		productController,
		transactionHistoryController,
		middlewareController,
	)

	if err := r.Run(":" + viper.GetString("PORT")); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}
