package product

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

func (r *Repository) CreateProduct(product model.Product) (model.Product, error) {
	err := r.DB.Create(&product).Error
	if err != nil {
		return model.Product{}, err
	}

	return product, nil
}

func (r *Repository) FetchListProducts() ([]model.Product, error) {
	data := []model.Product{}

	err := r.DB.Order("id ASC").Find(&data).Error
	if err != nil {
		return []model.Product{}, err
	}
	return data, nil
}

func (r *Repository) FetchProductByID(productID int) (model.Product, error) {
	var product model.Product

	err := r.DB.First(&product, productID).Error
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return model.Product{}, errors.New("product not found")
		default:
			return model.Product{}, err
		}
	}
	return product, nil
}

func (r *Repository) UpdateProduct(product model.Product) error {
	db := r.DB.Model(&model.Product{}).Where("id = ?", product.ID).Updates(product)
	err := db.Error
	if err != nil {
		return err
	}

	if db.RowsAffected < 1 {
		return errors.New("product not found")
	}
	return nil
}

func (r *Repository) DeleteProduct(productID int) error {
	return r.DB.Delete(&model.Product{}, productID).Error
}
