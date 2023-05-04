package services

import (
	"github.com/your_project/internal/api/models"
)

type CategoryRepository interface {
	GetAllCategories() ([]*models.Category, error)
	GetCategoryByID(id uint) (*models.Category, error)
}

type CategoryService struct {
	categoryRepo CategoryRepository
}

func NewCategoryService(categoryRepo CategoryRepository) *CategoryService {
	return &CategoryService{categoryRepo: categoryRepo}
}

func (s *CategoryService) GetAllCategories() ([]*models.Category, error) {
	return s.categoryRepo.GetAllCategories()
}

func (s *CategoryService) GetCategoryByID(id uint) (*models.Category, error) {
	return s.categoryRepo.GetCategoryByID(id)
}
