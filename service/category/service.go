package category

import (
	model "github.com/enricoanto/final-project/repository"
	categoryRepository "github.com/enricoanto/final-project/repository/category"
	"github.com/golang-jwt/jwt/v5"
)

type Service struct {
	categoryRepository *categoryRepository.Repository
}

type Claim struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func NewService(categoryRepository *categoryRepository.Repository) *Service {
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

func (s *Service) FetchCategoryByID(categoryID int) (model.Category, error) {
	return s.categoryRepository.FetchCategoryByID(categoryID)
}

func (s *Service) UpdateCategory(categoryID int, categoryType string) (model.Category, error) {
	err := s.categoryRepository.UpdateCategory(categoryID, categoryType)
	if err != nil {
		return model.Category{}, err
	}

	return s.FetchCategoryByID(categoryID)
}

func (s *Service) DeleteCategory(categoryID int) error {
	return s.categoryRepository.DeleteCategory(categoryID)
}
