package http

import (
	"net/http"

	"github.com/Yury132/Golang-Task-1/internal/transport/http/handlers"
	"github.com/gorilla/mux"
)

func InitRoutes(h *handlers.Handler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/health", h.Health).Methods(http.MethodGet)
	r.HandleFunc("/hello", h.Hello).Methods(http.MethodGet)

	return r
}
