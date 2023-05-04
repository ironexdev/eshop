package api

import (
	"e-shop-fiber/config"
	"e-shop-fiber/internal/api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type App struct {
	fiber *fiber.App
	cfg   *config.Config
}

func New(cfg *config.Config) *App {
	app := &App{
		fiber: fiber.New(),
		cfg:   cfg,
	}

	app.setupMiddlewares()
	app.setupStaticFiles()
	routes.SetupRoutes(app.fiber)

	return app
}

func (a *App) setupMiddlewares() {
	a.fiber.Use(recover.New())
}

func (a *App) setupStaticFiles() {
	a.fiber.Static("/static", "./static")
}

func (a *App) Start() {
	a.fiber.Listen(a.cfg.App.Address)
}
