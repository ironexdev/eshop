package repositories

import (
	"github.com/your_project/internal/api/models"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) GetAllProducts() ([]*models.Product, error) {
	var products []*models.Product
	err := r.db.Preload("Category").Find(&products).Error
	return products, err
}

func (r *ProductRepository) GetProductByID(id uint) (*models.Product, error) {
	var product models.Product
	err := r.db.Preload("Category").First(&product, id).Error
	return &product, err
}
