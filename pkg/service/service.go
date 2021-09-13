package service

import (
	"github.com/p12s/rshb-intex-refactoring/model"
	"github.com/p12s/rshb-intex-refactoring/pkg/repository"
)

type Book interface {
	GetBooksByAuthor(authorId int) ([]model.Book, error)
	GetAuthorBooksCount(authorId int) (int, error)
}

type Service struct {
	Book
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Book: NewBookService(repos.Book),
	}
}
