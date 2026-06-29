package book

import (
	"Bukus/internal/entity"
	"Bukus/internal/repo"
	"context"
)

type UseCase struct {
	repo   repo.BookRepo
	webAPI repo.BookWebApi
}

func New(r repo.BookRepo) *UseCase {
	return &UseCase{
		repo: r,
	}
}

func (obj *UseCase) Store(ctx context.Context, userID string, b entity.Book) error {
	return nil
}
