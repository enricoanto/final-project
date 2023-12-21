package product

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/enricoanto/final-project/helper"
	model "github.com/enricoanto/final-project/repository"
	"github.com/enricoanto/final-project/routes/middleware"

	productService "github.com/enricoanto/final-project/service/product"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	productService *productService.Service
}

func NewController(productService *productService.Service) *Controller {
	return &Controller{
		productService: productService,
	}
}

func (controller *Controller) CreateProduct(c *gin.Context) {
	claims, _ := middleware.Claims(c)

	role, _ := claims["role"].(string)
	if role != "admin" {
		helper.Error(c, http.StatusUnauthorized, errors.New("access denied"))
		return
	}

	var request ProductRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		helper.Error(c, http.StatusBadRequest, err)
		return
	}

	if err := helper.GetValidator().Struct(request); err != nil {
		helper.Error(c, http.StatusBadRequest, err)
		return
	}

	product := model.Product{
		Title:      request.Title,
		Price:      request.Price,
		Stock:      request.Stock,
		CategoryID: request.CategoryID,
	}

	product, err := controller.productService.CreateProduct(product)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err)
		return
	}
	product.UpdatedAt = time.Unix(0, 0)

	response := transformToProductResponse(product)

	helper.Success(c, http.StatusCreated, response)

}

func (controller *Controller) FetchListProducts(c *gin.Context) {
	products, err := controller.productService.FetchListProducts()
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err)
		return
	}

	response := transformToProductsResponse(products)

	helper.Success(c, http.StatusOK, response)
}

func (controller *Controller) UpdateProduct(c *gin.Context) {
	claims, _ := middleware.Claims(c)

	role, _ := claims["role"].(string)
	if role != "admin" {
		helper.Error(c, http.StatusUnauthorized, errors.New("access denied"))
		return
	}

	id, _ := strconv.Atoi(c.Param("productId"))
	var request ProductRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		helper.Error(c, http.StatusBadRequest, err)
		return
	}

	if err := helper.GetValidator().Struct(request); err != nil {
		helper.Error(c, http.StatusBadRequest, err)
		return
	}

	product := model.Product{
		ID:         id,
		Title:      request.Title,
		Price:      request.Price,
		Stock:      request.Stock,
		CategoryID: request.CategoryID,
	}

	product, err := controller.productService.UpdateProduct(product)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err)
		return
	}

	product.CreatedAt = time.Unix(0, 0)

	response := transformToProductResponse(product)

	helper.Success(c, http.StatusOK, response)
}

func (controller *Controller) DeleteProduct(c *gin.Context) {
	claims, _ := middleware.Claims(c)

	role, _ := claims["role"].(string)
	if role != "admin" {
		helper.Error(c, http.StatusUnauthorized, errors.New("access denied"))
		return
	}

	id, _ := strconv.Atoi(c.Param("productId"))

	err := controller.productService.DeleteProduct(id)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err)
		return
	}

	response := MessageResponse{
		Message: "Product has been successfully deleted",
	}

	helper.Success(c, http.StatusCreated, response)
}
