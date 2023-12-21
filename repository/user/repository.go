package user

import (
	"errors"

	"github.com/enricoanto/final-project/helper"
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

func (r *Repository) CreateUser(user model.User) (model.User, error) {
	err := r.DB.Create(&user).Error
	if err != nil {
		return model.User{}, err
	}

	return user, err
}

func (r *Repository) FindBy(user model.User) (model.User, error) {
	var data model.User

	err := r.DB.Where(&user).First(&data).Error
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return model.User{}, errors.New(helper.USER_NOT_FOUND)
		default:
			return model.User{}, err
		}
	}

	return data, nil
}

func (r *Repository) UpdateBalance(userID int, balance int) error {

	db := r.DB.Model(&model.User{}).Where("id = ?", userID).Update("balance", balance)
	err := db.Error
	if err != nil {
		return err
	}

	if db.RowsAffected < 1 {
		return errors.New(helper.USER_NOT_FOUND)
	}

	return nil
}
