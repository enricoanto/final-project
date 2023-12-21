package category

import (
	"errors"

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
			return model.Category{}, errors.New("category not found")
		default:
			return model.Category{}, err
		}
	}
	return category, nil
}

func (r *Repository) UpdateCategory(categoryID int, categoryType string) error {
	db := r.DB.Model(&model.Category{}).Where("id = ?", categoryID).Update("type", categoryType)
	err := db.Error
	if err != nil {
		return err
	}

	if db.RowsAffected < 1 {
		return errors.New("category not found")
	}
	return nil
}

func (r *Repository) DeleteCategory(categoryID int) error {
	return r.DB.Delete(&model.Category{}, categoryID).Error
}
