package restapi

import (
	"Bukus/config"
	"Bukus/internal/controller/restapi/middleware"
	v1 "Bukus/internal/controller/restapi/v1"
	"Bukus/internal/usecase"
	"Bukus/pkg/jwt"
	"Bukus/pkg/logger"
	"fmt"
	"github.com/gofiber/fiber/v3"
	"net/http"
)

func NewRouter(app *fiber.App, cfg *config.Config, uc *usecase.Container, jwtManager *jwt.Manager, l logger.Interface) {
	// Options
	app.Use(middleware.Logger(l))
	app.Use(middleware.Recovery(l))

	app.Get("/health", func(ctx fiber.Ctx) {
		ctx.Status(http.StatusOK)
	})

	// Routers
	apiV1Group := app.Group("/v1")
	{
		v1.NewRoutes(apiV1Group, uc, jwtManager, l)
	}

	for _, route := range app.GetRoutes() {
		fmt.Printf("%-6s %s\n", route.Method, route.Path)
	}
}
