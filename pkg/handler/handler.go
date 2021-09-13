package handler

import (
	"github.com/gorilla/mux"
	"github.com/p12s/rshb-intex-refactoring/pkg/service"
	"net/http"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/getBookCountByAuthor/{id}", h.getBookCountByAuthor).Methods(http.MethodGet)
	router.HandleFunc("/getBookByAuthor/{id}", h.getBookByAuthor).Methods(http.MethodGet)
	http.Handle("/", router)

	return router
}
