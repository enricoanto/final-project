package routes

import (
	"github.com/enricoanto/final-project/controller/category"
	"github.com/enricoanto/final-project/controller/product"
	"github.com/enricoanto/final-project/controller/transaction_history"
	"github.com/enricoanto/final-project/controller/user"
	"github.com/gin-gonic/gin"
)

func API(
	r *gin.Engine,
	userController *user.Controller,
	categoryController *category.Controller,
	productController *product.Controller,
	transactionHistoryController *transactionhistory.Controller,
	md *Controller,
) {

	user := r.Group("/users")
	user.POST("/register", userController.Register)
	user.POST("/login", userController.Login)
	user.PATCH("/topup", userController.UpdateBalance)

	category := r.Group("/categories", md.Middleware("admin"))
	category.POST("", categoryController.CreateCategory)
	category.GET("", categoryController.FetchListCategories)
	category.PATCH("/:categoryId", categoryController.UpdateCategory)
	category.DELETE("/:categoryId", categoryController.DeleteCategory)

	product := r.Group("/products")
	product.POST("", productController.CreateProduct)
	product.GET("", productController.FetchListProducts)
	product.PUT("/:productId", productController.UpdateProduct)
	product.DELETE("/:productId", productController.DeleteProduct)

	transactionHistory := r.Group("transactions")
	transactionHistory.POST("", transactionHistoryController.CreateTransactionHistory)
}
