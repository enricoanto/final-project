package routes

import (
	"github.com/enricoanto/final-project/controller/category"
	"github.com/enricoanto/final-project/controller/product"
	"github.com/enricoanto/final-project/controller/transaction_history"
	"github.com/enricoanto/final-project/controller/user"
	"github.com/enricoanto/final-project/routes/middleware"
	"github.com/gin-gonic/gin"
)

func API(
	r *gin.Engine,
	userController *user.Controller,
	categoryController *category.Controller,
	productController *product.Controller,
	transactionHistoryController *transactionhistory.Controller,
	md *middleware.Controller,
) {

	user := r.Group("/users")
	user.POST("/register", userController.Register)
	user.POST("/login", userController.Login)
	user.Use(md.Middleware())
	user.PATCH("/topup", userController.UpdateBalance)

	category := r.Group("/categories", md.Middleware())
	category.POST("", categoryController.CreateCategory)
	category.GET("", categoryController.FetchListCategories)
	category.PATCH("/:categoryId", categoryController.UpdateCategory)
	category.DELETE("/:categoryId", categoryController.DeleteCategory)

	product := r.Group("/products", md.Middleware())
	product.POST("", productController.CreateProduct)
	product.GET("", productController.FetchListProducts)
	product.PUT("/:productId", productController.UpdateProduct)
	product.DELETE("/:productId", productController.DeleteProduct)

	transactionHistory := r.Group("/transactions", md.Middleware())
	transactionHistory.POST("", transactionHistoryController.CreateTransactionHistory)
	transactionHistory.GET("/my-transactions", transactionHistoryController.MyTransaction)
	transactionHistory.GET("/user-transactions", transactionHistoryController.UserTransaction)
}
