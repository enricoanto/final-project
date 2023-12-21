package category

import (
	"encoding/json"
	"time"

	model "github.com/enricoanto/final-project/repository"
)

type (
	CategoryRequest struct {
		Type string `json:"type" validate:"required"`
	}

	CategoryResponse struct {
		ID                int        `json:"id"`
		Type              string     `json:"type"`
		SoldProductAmount int        `json:"sold_product_amount"`
		Products          []Product  `json:",omitempty"`
		CreatedAt         *time.Time `json:"created_at,omitempty"`
		UpdatedAt         *time.Time `json:"updated_at,omitempty"`
	}

	Product struct {
		ID        int       `json:"id"`
		Title     string    `json:"title"`
		Price     int       `json:"price"`
		Stock     int       `json:"stock"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	MessageResponse struct {
		Message string `json:"message"`
	}
)

func transformToCategoryResponse(data model.Category) CategoryResponse {
	response := CategoryResponse{
		ID:                data.ID,
		Type:              data.Type,
		SoldProductAmount: data.SoldProductAmount,
		Products:          transformProduct(data.Products),
	}

	if data.CreatedAt.Unix() != 0 {
		response.CreatedAt = &data.CreatedAt
	}

	if data.UpdatedAt.Unix() != 0 {
		response.UpdatedAt = &data.UpdatedAt
	}

	return response
}

func transformToCategoriesResponse(categories []model.Category) []CategoryResponse {
	response := []CategoryResponse{}

	for _, category := range categories {
		response = append(response, transformToCategoryResponse(category))
	}
	return response
}

func transformProduct(products []model.Product) []Product {
	response := []Product{}
	js, _ := json.Marshal(products)

	json.Unmarshal(js, &response)

	return response
}
