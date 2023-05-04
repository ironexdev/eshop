package models

import (
	"time"
)

type Order struct {
	ID          uint        `json:"id" gorm:"primaryKey"`
	UserID      uint        `json:"user_id"`
	User        User        `json:"user" gorm:"foreignKey:UserID"`
	TotalAmount float64     `json:"total_amount"`
	Status      OrderStatus `json:"status"`
	Items       []OrderItem `json:"items" gorm:"foreignKey:OrderID"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}

type OrderStatus string

const (
	OrderStatusCreated   OrderStatus = "created"
	OrderStatusProcessing OrderStatus = "processing"
	OrderStatusShipped    OrderStatus = "shipped"
	OrderStatusCompleted  OrderStatus = "completed"
	OrderStatusCancelled  OrderStatus = "cancelled"
)

type OrderItem struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	OrderID     uint      `json:"order_id"`
	ProductID   uint      `json:"product_id"`
	Product     Product   `json:"product" gorm:"foreignKey:ProductID"`
	Quantity    int       `json:"quantity"`
	Price       float64   `json:"price"`
	TotalAmount float64   `json:"total_amount"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
