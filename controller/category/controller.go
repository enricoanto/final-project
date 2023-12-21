package category

import (
	"net/http"
	"strconv"
	"time"

	"github.com/enricoanto/final-project/helper"
	model "github.com/enricoanto/final-project/repository"

	categoryService "github.com/enricoanto/final-project/service/category"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	categoryService *categoryService.Service
}

func NewController(categoryService *categoryService.Service) *Controller {
	return &Controller{
		categoryService: categoryService,
	}
}

func (controller *Controller) CreateCategory(c *gin.Context) {
	var request CategoryRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		helper.Error(c, http.StatusBadRequest, err)
		return
	}

	if err := helper.GetValidator().Struct(request); err != nil {
		helper.Error(c, http.StatusBadRequest, err)
		return
	}

	category := model.Category{
		Type: request.Type,
	}

	category, err := controller.categoryService.CreateCategory(category)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err)
		return
	}
	category.UpdatedAt = time.Unix(0, 0)

	response := transformToCategoryResponse(category)

	helper.Success(c, http.StatusCreated, response)
}

func (controller *Controller) FetchListCategories(c *gin.Context) {
	categories, err := controller.categoryService.FetchListCategories()
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err)
		return
	}

	response := transformToCategoriesResponse(categories)

	helper.Success(c, http.StatusOK, response)
}

func (controller *Controller) UpdateCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("categoryId"))
	var request CategoryRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		helper.Error(c, http.StatusBadRequest, err)
		return
	}

	if err := helper.GetValidator().Struct(request); err != nil {
		helper.Error(c, http.StatusBadRequest, err)
		return
	}

	category := model.Category{
		ID:   id,
		Type: request.Type,
	}

	category, err := controller.categoryService.UpdateCategory(category)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err)
		return
	}

	category.CreatedAt = time.Unix(0, 0)

	response := transformToCategoryResponse(category)

	helper.Success(c, http.StatusOK, response)
}

func (controller *Controller) DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("categoryId"))

	err := controller.categoryService.DeleteCategory(id)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err)
		return
	}

	response := MessageResponse{
		Message: "Category has been successfully deleted",
	}

	helper.Success(c, http.StatusCreated, response)
}
