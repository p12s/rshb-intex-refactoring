package repository

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/p12s/rshb-intex-refactoring/model"
)

type Book interface {
	GetBooksByAuthor(authorId int) ([]model.Book, error)
	GetAuthorBooksCount(authorId int) (int, error)
}

type Repository struct {
	Book
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		Book: NewBookPostgres(db),
	}
}
