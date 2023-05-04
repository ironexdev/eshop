package controllers

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/your_project/internal/api/services"
)

type CartController struct {
	cartService services.CartService
}

func NewCartController(cartService services.CartService) *CartController {
	return &CartController{cartService: cartService}
}

func (c *CartController) GetCart(ctx *fiber.Ctx) error {
	userID := ctx.Locals("userID").(uint)
	cart, err := c.cartService.GetCartByUserID(userID)
	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Cart not found"})
	}
	return ctx.Status(http.StatusOK).JSON(cart)
}

func (c *CartController) AddProduct(ctx *fiber.Ctx) error {
	userID := ctx.Locals("userID").(uint)
	productID, _ := strconv.Atoi(ctx.Params("productID"))
	quantity := uint(ctx.Query("quantity", "1"))

	err := c.cartService.AddProductToCart(userID, uint(productID), quantity)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to add product to cart"})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": "Product added to cart"})
}

func (c *CartController) RemoveProduct(ctx *fiber.Ctx) error {
	userID := ctx.Locals("userID").(uint)
	productID, _ := strconv.Atoi(ctx.Params("productID"))

	err := c.cartService.RemoveProductFromCart(userID, uint(productID))
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to remove product from cart"})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": "Product removed from cart"})
}

func (c *CartController) UpdateProductQuantity(ctx *fiber.Ctx) error {
	userID := ctx.Locals("userID").(uint)
	productID, _ := strconv.Atoi(ctx.Params("productID"))
	quantity := uint(ctx.Query("quantity", "1"))

	err := c.cartService.UpdateProductQuantity(userID, uint(productID), quantity)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update product quantity"})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": "Product quantity updated"})
}
