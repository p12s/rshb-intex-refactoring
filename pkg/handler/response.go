package handler

import (
	"encoding/json"
	"net/http"

	"github.com/p12s/rshb-intex-refactoring/model"
)

func NewErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	http.Error(w, message, statusCode)
}

func NewOkResponse(w http.ResponseWriter, message map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(message)
}

func NewBookResponse(w http.ResponseWriter, books []model.Book) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(books)
}
