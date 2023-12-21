package category

import (
	model "github.com/enricoanto/final-project/repository"
	"github.com/enricoanto/final-project/repository/category"
)

type Service struct {
	categoryRepository *category.Repository
}

func NewService(categoryRepository *category.Repository) *Service {
	return &Service{
		categoryRepository: categoryRepository,
	}
}

func (s *Service) CreateCategory(category model.Category) (model.Category, error) {
	return s.categoryRepository.CreateCategory(category)
}

func (s *Service) FetchListCategories() ([]model.Category, error) {
	return s.categoryRepository.FetchListCategories()
}

func (s *Service) UpdateCategory(category model.Category) (model.Category, error) {
	err := s.categoryRepository.UpdateCategory(category)
	if err != nil {
		return model.Category{}, err
	}

	return s.categoryRepository.FetchCategoryByID(category.ID)
}

func (s *Service) DeleteCategory(categoryID int) error {
	return s.categoryRepository.DeleteCategory(categoryID)
}

func (s *Service) FetchCategoryByID(categoryID int) (model.Category, error) {
	return s.categoryRepository.FetchCategoryByID(categoryID)
}
