package controllers

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/your_project/internal/api/services"
)

type CategoryController struct {
	categoryService services.CategoryService
}

func NewCategoryController(categoryService services.CategoryService) *CategoryController {
	return &CategoryController{categoryService: categoryService}
}

func (c *CategoryController) GetAllCategories(ctx *fiber.Ctx) error {
	categories, err := c.categoryService.GetAllCategories()
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch categories"})
	}
	return ctx.Status(http.StatusOK).JSON(categories)
}

func (c *CategoryController) GetCategoryByID(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	category, err := c.categoryService.GetCategoryByID(uint(id))
	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Category not found"})
	}
	return ctx.Status(http.StatusOK).JSON(category)
}
