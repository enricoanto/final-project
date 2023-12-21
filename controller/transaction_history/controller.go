package transactionhistory

import (
	"net/http"
	// "strconv"
	// "time"

	"github.com/enricoanto/final-project/helper"
	model "github.com/enricoanto/final-project/repository"

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
		UserID:    1,
	}

	transactionHistory, err := controller.transactionHistoryService.CreateTransactionHistory(transactionHistory)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err)
		return
	}

	response := transformToCreateResponse(transactionHistory)

	helper.Success(c, http.StatusCreated, response)
}
