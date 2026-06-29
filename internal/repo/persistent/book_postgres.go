package persistent

import (
	"Bukus/internal/entity"
	"Bukus/pkg/postgres"
	"context"
)

const _defaultEntityCap = 64

type BookRepo struct {
	*postgres.Postgres
}

func NewBookRepo(pg *postgres.Postgres) *BookRepo {
	return &BookRepo{pg}
}

func (r *BookRepo) Store(ctx context.Context, userID string, t entity.Book) error {
	//sql, args, err := r.Builder.
	//	Insert("history").
	//	Columns("user_id, source, destination, original, translation").
	//	Values(userID, b.Source, b.Destination, b.Original, b.Translation).
	//	ToSql()
	//if err != nil {
	//	return fmt.Errorf("TranslationRepo - Store - r.Builder: %w", err)
	//}
	//
	//_, err = r.Pool.Exec(ctx, sql, args...)
	//if err != nil {
	//	return fmt.Errorf("TranslationRepo - Store - r.Pool.Exec: %w", err)
	//}

	return nil
}
