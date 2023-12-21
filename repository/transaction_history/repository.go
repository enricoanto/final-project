package transactionHistory

import (
	// "errors"

	model "github.com/enricoanto/final-project/repository"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(DB *gorm.DB) *Repository {
	return &Repository{
		DB: DB,
	}
}

func (r *Repository) CreateTransactionHistory(transactionHistory model.TransactionHistory) (model.TransactionHistory, error) {
	err := r.DB.Create(&transactionHistory).Error
	if err != nil {
		return model.TransactionHistory{}, err
	}

	return transactionHistory, nil
}

func (r *Repository) FetchListTransactionHistories(transactionHistory model.TransactionHistory) ([]model.TransactionHistory, error) {
	var data []model.TransactionHistory

	err := r.DB.Preload("User").Preload("Product").Find(&data).Error
	if err != nil {
		return []model.TransactionHistory{}, err
	}

	return data, nil
}
