package product

import (
	model "github.com/enricoanto/final-project/repository"
	productRepository "github.com/enricoanto/final-project/repository/product"
)

type Service struct {
	productRepository *productRepository.Repository
}

func NewService(productRepository *productRepository.Repository) *Service {
	return &Service{
		productRepository: productRepository,
	}
}

func (s *Service) CreateProduct(product model.Product) (model.Product, error) {
	return s.productRepository.CreateProduct(product)
}

func (s *Service) FetchListProducts() ([]model.Product, error) {
	return s.productRepository.FetchListProducts()
}

func (s *Service) UpdateProduct(product model.Product) (model.Product, error) {
	err := s.productRepository.UpdateProduct(product)
	if err != nil {
		return model.Product{}, err
	}

	return s.productRepository.FetchProductByID(product.ID)
}

func (s *Service) DeleteProduct(productID int) error {
	return s.productRepository.DeleteProduct(productID)
}
