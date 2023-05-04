package services

import (
	"github.com/your_project/internal/api/models"
)

type ProductRepository interface {
	GetAllProducts() ([]*models.Product, error)
	GetProductByID(id uint) (*models.Product, error)
}

type ProductService struct {
	productRepo ProductRepository
}

func NewProductService(productRepo ProductRepository) *ProductService {
	return &ProductService{productRepo: productRepo}
}

func (s *ProductService) GetAllProducts() ([]*models.Product, error) {
	return s.productRepo.GetAllProducts()
}

func (s *ProductService) GetProductByID(id uint) (*models.Product, error) {
	return s.productRepo.GetProductByID(id)
}
