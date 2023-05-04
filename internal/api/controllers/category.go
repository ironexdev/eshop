package controllers

import (
	"e-shop-fiber/internal/api/models"
	"github.com/gofiber/fiber/v2"
)

type CategoryController struct {
	categoryModel *models.CategoryModel
}

func NewCategoryController() *CategoryController {
	return &CategoryController{
		categoryModel: models.NewCategoryModel(),
	}
}

func (c *CategoryController) List(c *fiber.Ctx) error {
	// Implement listing categories logic here
	return c.JSON(fiber.Map{"message": "Categories listed"})
}
