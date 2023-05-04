package services

import (
	"errors"

	"github.com/your_project/internal/models"
)

type CartRepository interface {
	GetCartByUserID(userID uint) (*models.Cart, error)
	AddProductToCart(cart *models.Cart, productID uint, quantity uint) error
	RemoveProductFromCart(cart *models.Cart, productID uint) error
	UpdateProductQuantity(cart *models.Cart, productID uint, quantity uint) error
}

type CartService struct {
	cartRepo CartRepository
}

func NewCartService(cartRepo CartRepository) *CartService {
	return &CartService{
		cartRepo: cartRepo,
	}
}

func (s *CartService) GetCartByUserID(userID uint) (*models.Cart, error) {
	return s.cartRepo.GetCartByUserID(userID)
}

func (s *CartService) AddProductToCart(userID uint, productID uint, quantity uint) error {
	cart, err := s.cartRepo.GetCartByUserID(userID)
	if err != nil {
		return errors.New("Cart not found")
	}
	return s.cartRepo.AddProductToCart(cart, productID, quantity)
}

func (s *CartService) RemoveProductFromCart(userID uint, productID uint) error {
	cart, err := s.cartRepo.GetCartByUserID(userID)
	if err != nil {
		return errors.New("Cart not found")
	}
	return s.cartRepo.RemoveProductFromCart(cart, productID)
}

func (s *CartService) UpdateProductQuantity(userID uint, productID uint, quantity uint) error {
	cart, err := s.cartRepo.GetCartByUserID(userID)
	if err != nil {
		return errors.New("Cart not found")
	}
	return s.cartRepo.UpdateProductQuantity(cart, productID, quantity)
}
