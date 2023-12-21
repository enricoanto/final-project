package transactionhistory

import (
	"errors"
	"net/http"

	"github.com/enricoanto/final-project/helper"
	model "github.com/enricoanto/final-project/repository"

	"github.com/enricoanto/final-project/routes/middleware"
	"github.com/enricoanto/final-project/service/transaction_history"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	transactionHistoryService *transactionhistory.Service
}

func NewController(transactionHistoryService *transactionhistory.Service) *Controller {
	return &Controller{
		transactionHistoryService: transactionHistoryService,
	}
}

func (controller *Controller) CreateTransactionHistory(c *gin.Context) {
	claims, _ := middleware.Claims(c)

	userID, _ := claims["id"].(float64)

	var request TransactionHistoryRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		helper.Error(c, http.StatusBadRequest, err)
		return
	}

	if err := helper.GetValidator().Struct(request); err != nil {
		helper.Error(c, http.StatusBadRequest, err)
		return
	}

	transactionHistory := model.TransactionHistory{
		ProductID: request.ProductID,
		Quantity:  request.Quantity,
		UserID:    int(userID),
	}

	transactionHistory, err := controller.transactionHistoryService.CreateTransactionHistory(transactionHistory)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err)
		return
	}

	response := transformToCreateResponse(transactionHistory)

	helper.Success(c, http.StatusCreated, response)
}

func (controller *Controller) MyTransaction(c *gin.Context) {
	claims, _ := middleware.Claims(c)

	userID, _ := claims["id"].(float64)

	filter := model.TransactionHistory{
		UserID: int(userID),
	}

	transactions, err := controller.transactionHistoryService.FetchListTransactionHistories(filter)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err)
		return
	}

	response := transformToMyTransactions(transactions)

	helper.Success(c, http.StatusOK, response)
}

func (controller *Controller) UserTransaction(c *gin.Context) {
	claims, _ := middleware.Claims(c)

	role, _ := claims["role"].(string)
	if role != "admin" {
		helper.Error(c, http.StatusUnauthorized, errors.New("access denied"))
		return
	}

	transactions, err := controller.transactionHistoryService.FetchListTransactionHistories(model.TransactionHistory{})
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err)
		return
	}

	response := transformToUserTransactions(transactions)
	helper.Success(c, http.StatusOK, response)
}
