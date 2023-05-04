package services

import (
	"errors"

	"github.com/your_project/internal/api/models"
)

type OrderRepository interface {
	CreateOrder(order *models.Order) error
	GetOrderByID(id uint) (*models.Order, error)
	GetOrdersByUserID(userID uint) ([]*models.Order, error)
}

type OrderService struct {
	orderRepo OrderRepository
}

func NewOrderService(orderRepo OrderRepository) *OrderService {
	return &OrderService{orderRepo: orderRepo}
}

func (s *OrderService) CreateOrder(order *models.Order) error {
	return s.orderRepo.CreateOrder(order)
}

func (s *OrderService) GetOrderByID(id uint) (*models.Order, error) {
	return s.orderRepo.GetOrderByID(id)
}

func (s *OrderService) GetOrdersByUserID(userID uint) ([]*models.Order, error) {
	return s.orderRepo.GetOrdersByUserID(userID)
}
