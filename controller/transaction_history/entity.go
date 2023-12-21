package transactionhistory

import (
	"encoding/json"
	"time"

	model "github.com/enricoanto/final-project/repository"
)

type (
	TransactionHistoryRequest struct {
		ProductID int `json:"product_id" validate:"required"`
		Quantity  int `json:"quantity" validate:"required"`
	}

	CreateResponse struct {
		Message        string         `json:"message"`
		TransctionBill TransctionBill `json:"transaction_bill"`
	}

	TransctionBill struct {
		TotalPrice   int    `json:"total_price"`
		Quantity     int    `json:"quantity"`
		ProductTitle string `json:"product_title"`
	}
	MyTransaction struct {
		ID         int     `json:"id"`
		ProductID  int     `json:"product_id"`
		UserID     int     `json:"user_id"`
		Quantity   int     `json:"quantity"`
		TotalPrice int     `json:"total_price"`
		Product    Product `json:"Product"`
	}
	UserTransaction struct {
		MyTransaction
		User User
	}

	User struct {
		ID        int       `json:"id"`
		FullName  string    `json:"full_name"`
		Email     string    `json:"email"`
		Balance   int       `json:"balance"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	Product struct {
		ID         int       `json:"id"`
		Title      string    `json:"title"`
		Price      int       `json:"price"`
		Stock      int       `json:"stock"`
		CategoryID int       `json:"category_Id"`
		CreatedAt  time.Time `json:"created_at"`
		UpdatedAt  time.Time `json:"updated_at"`
	}
)

func transformToCreateResponse(data model.TransactionHistory) CreateResponse {
	return CreateResponse{
		Message: "You have succesfully purchased the product",
		TransctionBill: TransctionBill{
			TotalPrice:   data.TotalPrice,
			Quantity:     data.Quantity,
			ProductTitle: data.Product.Title,
		},
	}
}
func transformToMyTransactions(data []model.TransactionHistory) []MyTransaction {
	response := []MyTransaction{}

	js, _ := json.Marshal(data)
	json.Unmarshal(js, &response)

	return response
}

func transformToUserTransactions(data []model.TransactionHistory) []UserTransaction {
	response := []UserTransaction{}

	js, _ := json.Marshal(data)
	json.Unmarshal(js, &response)

	return response
}
