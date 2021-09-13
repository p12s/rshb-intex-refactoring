package handler

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// Получение списка книг автора по его Id
func (h *Handler) getBookByAuthor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		NewErrorResponse(w, http.StatusBadRequest, "Error: empty `id` param!")
		return
	}

	authorId, err := strconv.Atoi(id)
	if err != nil {
		NewErrorResponse(w, http.StatusBadRequest, "Error: can't convert `id` param to int!")
		return
	}

	books, err := h.services.GetBooksByAuthor(authorId)
	if err != nil {
		logrus.Fatalf("Error during method execution getBookByAuthor-GetBooksByAuthor: %s\n", err.Error())
		NewErrorResponse(w, http.StatusInternalServerError, "Error during method execution GetBooksByAuthor")
		return
	}

	NewBookResponse(w, books)
}

// Получение кол-ва книг автора по его Id
func (h *Handler) getBookCountByAuthor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		NewErrorResponse(w, http.StatusBadRequest, "Error: empty `id` param!")
		return
	}

	authorId, err := strconv.Atoi(id)
	if err != nil {
		NewErrorResponse(w, http.StatusBadRequest, "Error: can't convert `id` param to int!")
		return
	}

	booksCount, err := h.services.GetAuthorBooksCount(authorId)
	if err != nil {
		logrus.Fatalf("Error during method execution getBookCountByAuthor-GetAuthorBooksCount: %s\n", err.Error())
		NewErrorResponse(w, http.StatusInternalServerError, "Error during method execution GetAuthorBooksCount")
		return
	}

	NewOkResponse(w, map[string]interface{}{
		"count": booksCount,
	})
}
