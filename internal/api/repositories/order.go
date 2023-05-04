package repositories

import (
	"github.com/your_project/internal/api/models"
	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) CreateOrder(order *models.Order) error {
	return r.db.Create(&order).Error
}

func (r *OrderRepository) GetOrderByID(id uint) (*models.Order, error) {
	var order models.Order
	err := r.db.Preload("OrderItems.Product").First(&order, id).Error
	return &order, err
}

func (r *OrderRepository) GetOrdersByUserID(userID uint) ([]*models.Order, error) {
	var orders []*models.Order
	err := r.db.Where("user_id = ?", userID).Preload("OrderItems.Product").Find(&orders).Error
	return orders, err
}
