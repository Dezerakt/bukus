package persistent

import (
	"bukus/internal/entity"
	"bukus/pkg/postgres"
	"context"
)

const _defaultEntityCap = 64

type BookRepo struct {
	*postgres.Postgres
}

func NewBookRepo(pg *postgres.Postgres) *BookRepo {
	return &BookRepo{pg}
}

func (obj *BookRepo) Store(ctx context.Context, book entity.Book) error {

	return nil
}
