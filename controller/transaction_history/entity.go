package transactionhistory

import model "github.com/enricoanto/final-project/repository"

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
)

func transformToCreateResponse(data model.TransactionHistory) CreateResponse {
	return CreateResponse{
		Message: "You have succesfully purchased the product",
		TransctionBill: TransctionBill{
			TotalPrice: data.TotalPrice,
			Quantity: data.Quantity,
			ProductTitle: data.Product.Title,
		},
	}
}
