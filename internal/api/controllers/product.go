package controllers

import (
	"e-shop-fiber/internal/api/models"
	"github.com/gofiber/fiber/v2"
)

type ProductController struct {
	productModel *models.ProductModel
}

func NewProductController() *ProductController {
	return &ProductController{
		productModel: models.NewProductModel(),
	}
}

func (p *ProductController) List(c *fiber.Ctx) error {
	// Implement listing products logic here
	return c.JSON(fiber.Map{"message": "Products listed"})
}

func (p *ProductController) Get(c *fiber.Ctx) error {
	// Implement getting single product details logic here
	return c.JSON(fiber.Map{"message": "Product details fetched"})
}
