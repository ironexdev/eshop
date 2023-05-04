package routes

import (
	"e-shop-fiber/internal/api/controllers"
	"e-shop-fiber/internal/api/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	authController := controllers.NewAuthController()
	api.Post("/register", authController.Register)
	api.Post("/login", authController.Login)
	api.Post("/logout", authController.Logout)

	productController := controllers.NewProductController()
	api.Get("/products", productController.List)
	api.Get("/products/:id", productController.Get)

	authenticated := api.Use(middlewares.AuthRequired())
	cartController := controllers.NewCartController()
	authenticated.Post("/cart", cartController.Add)
	authenticated.Get("/cart", cartController.List)
	authenticated.Delete("/cart/:id", cartController.Remove)

	orderController := controllers.NewOrderController()
	authenticated.Post("/orders", orderController.Create)
	authenticated.Get("/orders", orderController.List)
	authenticated.Get("/orders/:id", orderController.Get)
}
