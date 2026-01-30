package book

import (
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux, handler *Handler) {
	mux.HandleFunc("POST /book", Save(handler.store))
	mux.HandleFunc("GET /book", Filter(handler.store))
	mux.HandleFunc("GET /book/{id}", Get(handler.store))
	mux.HandleFunc("PUT /book/{id}", Update(handler.store))
	mux.HandleFunc("DELETE /book/{id}", Delete(handler.store))
}
