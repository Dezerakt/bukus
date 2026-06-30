package usecase

import (
	"bukus/internal/entity"
	"context"
)

type (
	Book interface {
		Store(ctx context.Context, b entity.Book) error
	}
)
