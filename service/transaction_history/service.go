package transactionhistory

import (
	"errors"

	model "github.com/enricoanto/final-project/repository"
	"github.com/enricoanto/final-project/repository/transaction_history"
	"github.com/enricoanto/final-project/service/category"
	"github.com/enricoanto/final-project/service/product"
	"github.com/enricoanto/final-project/service/user"
)

type Service struct {
	transactionHistoryRepository *transactionHistory.Repository
	categoryService              *category.Service
	productService               *product.Service
	userService                  *user.Service
}

func NewService(
	transactionHistoryRepository *transactionHistory.Repository,
	categoryService *category.Service,
	productService *product.Service,
	userService *user.Service,
) *Service {
	return &Service{
		transactionHistoryRepository,
		categoryService,
		productService,
		userService,
	}
}

func (s *Service) CreateTransactionHistory(transactionHistory model.TransactionHistory) (model.TransactionHistory, error) {
	product, err := s.productService.FetchProductByID(transactionHistory.ProductID)
	if err != nil {
		return model.TransactionHistory{}, err
	}

	user, err := s.userService.FindBy(model.User{ID: transactionHistory.UserID})
	if err != nil {
		return model.TransactionHistory{}, err
	}

	product.Stock = product.Stock - transactionHistory.Quantity
	if product.Stock < 0 {
		return model.TransactionHistory{}, errors.New("product out of stock")
	}

	transactionHistory.TotalPrice = product.Price * transactionHistory.Quantity

	purchased := user.Balance - transactionHistory.TotalPrice

	if purchased < 0 {
		return model.TransactionHistory{}, errors.New("balance insufficient")
	}

	_, err = s.userService.UpdateBalance(transactionHistory.UserID, -transactionHistory.TotalPrice)

	product, err = s.productService.UpdateProduct(product)
	if err != nil {
		return model.TransactionHistory{}, err
	}

	category, err := s.categoryService.FetchCategoryByID(product.CategoryID)
	if err != nil {
		return model.TransactionHistory{}, err
	}

	category.SoldProductAmount = category.SoldProductAmount + transactionHistory.Quantity

	_, err = s.categoryService.UpdateCategory(category)
	if err != nil {
		return model.TransactionHistory{}, err
	}

	transHistory, err := s.transactionHistoryRepository.CreateTransactionHistory(transactionHistory)
	if err != nil {
		return model.TransactionHistory{}, err
	}
	transHistory.Product = product

	return transHistory, nil
}

func (s *Service) FetchListTransactionHistories(transactionHistory model.TransactionHistory) ([]model.TransactionHistory, error) {
	return s.transactionHistoryRepository.FetchListTransactionHistories(transactionHistory)
}
