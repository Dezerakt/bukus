package repo

import (
	"bukus/internal/entity"
	"context"
)

type (
	BookRepo interface {
		Store(ctx context.Context, book entity.Book) error
	}

	BookWebApi interface {
	}
)
