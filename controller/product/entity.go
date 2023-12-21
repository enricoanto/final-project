package product

import (
	"time"

	model "github.com/enricoanto/final-project/repository"
)

type (
	ProductRequest struct {
		Title      string `json:"title" validate:"required"`
		Price      int    `json:"price" validate:"required,min=0,max=50000000"`
		Stock      int    `json:"stock" validate:"required,min=5"`
		CategoryID int    `json:"category_Id" validate:"required"`
	}

	ProductResponse struct {
		ID         int        `json:"id"`
		Title      string     `json:"title"`
		Price      int        `json:"price"`
		Stock      int        `json:"stock"`
		CategoryID int        `json:"category_Id"`
		CreatedAt  *time.Time `json:"created_at,omitempty"`
		UpdatedAt  *time.Time `json:"updated_at,omitempty"`
	}

	MessageResponse struct {
		Message string `json:"message"`
	}
)

func transformToProductResponse(product model.Product) ProductResponse {
	response := ProductResponse{
		ID:         product.ID,
		Title:      product.Title,
		Price:      product.Price,
		Stock:      product.Stock,
		CategoryID: product.CategoryID,
	}

	if product.CreatedAt.Unix() != 0 {
		response.CreatedAt = &product.CreatedAt
	}

	if product.UpdatedAt.Unix() != 0 {
		response.UpdatedAt = &product.UpdatedAt
	}

	return response
}

func transformToProductsResponse(products []model.Product) []ProductResponse {
	response := []ProductResponse{}

	for _, product := range products {
		product.UpdatedAt = time.Unix(0, 0)
		response = append(response, transformToProductResponse(product))
	}
	return response
}
