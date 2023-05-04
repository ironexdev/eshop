package controllers

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/your_project/internal/api/models"
	"github.com/your_project/internal/api/services"
)

type OrderController struct {
	orderService services.OrderService
}

func NewOrderController(orderService services.OrderService) *OrderController {
	return &OrderController{orderService: orderService}
}

func (c *OrderController) CreateOrder(ctx *fiber.Ctx) error {
	order := &models.Order{}
	if err := ctx.BodyParser(order); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	userID := ctx.Locals("userID").(uint)
	order.UserID = userID

	if err := c.orderService.CreateOrder(order); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create order"})
	}

	return ctx.Status(http.StatusCreated).JSON(order)
}

func (c *OrderController) GetOrder(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	order, err := c.orderService.GetOrderByID(uint(id))
	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Order not found"})
	}

	return ctx.Status(http.StatusOK).JSON(order)
}

func (c *OrderController) GetOrders(ctx *fiber.Ctx) error {
	userID := ctx.Locals("userID").(uint)
	orders, err := c.orderService.GetOrdersByUserID(userID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch orders"})
	}

	return ctx.Status(http.StatusOK).JSON(orders)
}
