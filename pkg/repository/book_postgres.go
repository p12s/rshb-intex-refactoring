package repository

import (
	"context"
	"fmt"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/p12s/rshb-intex-refactoring/model"
)

type BookPostgres struct {
	db *pgxpool.Pool
}

func NewBookPostgres(db *pgxpool.Pool) *BookPostgres {
	return &BookPostgres{db: db}
}

func (r *BookPostgres) GetBooksByAuthor(authorId int) ([]model.Book, error) {
	query := fmt.Sprintf(`SELECT b.id, b.title, b.author_id, b.cost FROM %s b INNER JOIN %s a on a.id = b.author_id
		WHERE a.id = $1`,
		bookTable, authorTable)

	var items []model.Book
	err := pgxscan.Select(context.Background(), r.db, &items, query, authorId)
	return items, err
}

func (r *BookPostgres) GetAuthorBooksCount(authorId int) (int, error) {
	var bookCount int
	query := fmt.Sprintf(`SELECT COUNT(*) FROM %s b INNER JOIN %s a on a.id = b.author_id
		WHERE a.id = $1`, bookTable, authorTable)
	err := pgxscan.Get(context.Background(), r.db, &bookCount, query, authorId)
	return bookCount, err
}
