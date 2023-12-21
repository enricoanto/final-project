package category

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

func (r *Repository) CreateCategory(category model.Category) (model.Category, error) {
	err := r.DB.Create(&category).Error
	if err != nil {
		return model.Category{}, err
	}

	return category, nil
}

func (r *Repository) FetchListCategories() ([]model.Category, error) {
	data := []model.Category{}

	err := r.DB.Preload("Products", func(tx *gorm.DB) *gorm.DB {
		return tx.Order("id ASC")
	}).Order("id ASC").Find(&data).Error
	if err != nil {
		return []model.Category{}, err
	}
	return data, nil
}

func (r *Repository) FetchCategoryByID(categoryID int) (model.Category, error) {
	var category model.Category

	err := r.DB.First(&category, categoryID).Error
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return model.Category{}, errors.New(helper.CATEGORY_NOT_FOUND)
		default:
			return model.Category{}, err
		}
	}
	return category, nil
}

func (r *Repository) UpdateCategory(category model.Category) error {
	db := r.DB.Model(&model.Category{}).Where("id = ?", category.ID).Updates(category)
	err := db.Error
	if err != nil {
		return err
	}

	if db.RowsAffected < 1 {
		return errors.New(helper.CATEGORY_NOT_FOUND)
	}
	return nil
}

func (r *Repository) DeleteCategory(categoryID int) error {
	return r.DB.Delete(&model.Category{}, categoryID).Error
}
