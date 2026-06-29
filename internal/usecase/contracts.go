package usecase

import (
	"Bukus/internal/entity"
	"context"
)

type (
	Book interface {
		Store(ctx context.Context, userID string, b entity.Book) error
	}
)
