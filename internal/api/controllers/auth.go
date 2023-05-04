package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/your_project/internal/models"
	"github.com/your_project/internal/api/services"
)

type AuthController struct {
	authService services.AuthService
}

func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (a *AuthController) Register(ctx *fiber.Ctx) error {
	var user models.User
	err := ctx.BodyParser(&user)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Bad request"})
	}

	err = a.authService.Register(&user)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to register user"})
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{"message": "User registered successfully"})
}

func (a *AuthController) Login(ctx *fiber.Ctx) error {
	type LoginInput struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	input := new(LoginInput)
	if err := ctx.BodyParser(input); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Bad request"})
	}

	user, err := a.authService.Login(input.Email, input.Password)
	if err != nil {
		return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	token, err := a.authService.GenerateToken(user)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate token"})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": "Logged in successfully", "user": user, "token": token})
}
