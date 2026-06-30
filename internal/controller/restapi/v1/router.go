package v1

import (
	"bukus/internal/usecase"
	"bukus/pkg/jwt"
	"bukus/pkg/logger"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

// NewRoutes -.
func NewRoutes(apiV1Group fiber.Router, uc *usecase.Container, jwtManager *jwt.Manager, l logger.Interface) {
	r := &V1{
		bookUC: uc.Book,
		l:      l,
		v:      validator.New(validator.WithRequiredStructEnabled()),
	}

	// Public routes
	//authGroup := apiV1Group.Group("/auth")
	//{
	//	authGroup.Post("/register", r.register)
	//	authGroup.Post("/login", r.login)
	//}

	apiPrefix := apiV1Group.Group("/api")
	bookGroup := apiPrefix.Group("/book")
	{
		bookGroup.Post("/store", r.store)
	}

	// Protected routes
	//protected := apiV1Group.Group("", middleware.Auth(jwtManager))
	//
	//translationGroup := protected.Group("/book")
	//{
	//	translationGroup.Get("/history", r.history)
	//	translationGroup.Post("/do-translate", r.doTranslate)
	//}
}
