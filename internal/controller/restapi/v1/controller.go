package v1

import (
	"bukus/internal/usecase"
	"bukus/pkg/logger"

	"github.com/go-playground/validator/v10"
)

// V1 -.
type V1 struct {
	bookUC usecase.Book
	l      logger.Interface
	v      *validator.Validate
}
