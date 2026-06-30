package book

import (
	"bukus/internal/entity"
	"bukus/internal/repo"
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

func (obj *UseCase) Store(ctx context.Context, b entity.Book) error {

	return nil
}
