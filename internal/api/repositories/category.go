package repositories

import (
	"github.com/your_project/internal/api/models"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) GetAllCategories() ([]*models.Category, error) {
	var categories []*models.Category
	err := r.db.Find(&categories).Error
	return categories, err
}

func (r *CategoryRepository) GetCategoryByID(id uint) (*models.Category, error) {
	var category models.Category
	err := r.db.First(&category, id).Error
	return &category, err
}
