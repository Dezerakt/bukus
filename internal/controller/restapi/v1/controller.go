package v1

import (
	"Bukus/internal/usecase"
	"Bukus/pkg/logger"
	"github.com/go-playground/validator/v10"
)

// V1 -.
type V1 struct {
	uc *usecase.Container
	l  logger.Interface
	v  *validator.Validate
}
