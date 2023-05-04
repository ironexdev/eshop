package controllers

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/your_project/internal/api/services"
)

type ProductController struct {
	productService services.ProductService
}

func NewProductController(productService services.ProductService) *ProductController {
	return &ProductController{productService: productService}
}

func (c *ProductController) GetAllProducts(ctx *fiber.Ctx) error {
	products, err := c.productService.GetAllProducts()
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch products"})
	}
	return ctx.Status(http.StatusOK).JSON(products)
}

func (c *ProductController) GetProductByID(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	product, err := c.productService.GetProductByID(uint(id))
	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Product not found"})
	}
	return ctx.Status(http.StatusOK).JSON(product)
}
