package routes

import (
	"github.com/enricoanto/final-project/controller/category"
	"github.com/enricoanto/final-project/controller/user"
	"github.com/gin-gonic/gin"
)

func API(
	r *gin.Engine,
	userController *user.Controller,
	categoryController *category.Controller,
) {
	user := r.Group("/users")
	user.POST("/register", userController.Register)
	user.POST("/login", userController.Login)
	user.PATCH("/topup", userController.UpdateBalance)

	category := r.Group("categories")
	category.POST("", categoryController.CreateCategory)
	category.GET("", categoryController.FetchListCategories)
	category.PATCH("/:categoryId", categoryController.UpdateCategory)
	category.DELETE("/:categoryId", categoryController.DeleteCategory)
}