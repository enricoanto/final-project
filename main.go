package main

import (
	"github.com/enricoanto/final-project/config"
	"github.com/enricoanto/final-project/routes"
	"github.com/gin-gonic/gin"
	"log"

	// Controller
	categoryController "github.com/enricoanto/final-project/controller/category"
	productController "github.com/enricoanto/final-project/controller/product"
	userController "github.com/enricoanto/final-project/controller/user"

	// Repository
	categoryRepository "github.com/enricoanto/final-project/repository/category"
	productRepository "github.com/enricoanto/final-project/repository/product"
	userRepository "github.com/enricoanto/final-project/repository/user"

	// Service
	categoryService "github.com/enricoanto/final-project/service/category"
	productService "github.com/enricoanto/final-project/service/product"
	userService "github.com/enricoanto/final-project/service/user"
)

func main() {
	r := gin.Default()
	DB := config.DB

	userRepository := userRepository.NewRepository(DB)
	userService := userService.NewService(userRepository)
	userController := userController.NewController(userService)

	categoryRepository := categoryRepository.NewRepository(DB)
	categoryService := categoryService.NewService(categoryRepository)
	categoryController := categoryController.NewController(categoryService)

	productRepository := productRepository.NewRepository(DB)
	productService := productService.NewService(productRepository)
	productController := productController.NewController(productService)

	routes.API(
		r,
		userController,
		categoryController,
		productController,
	)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}
