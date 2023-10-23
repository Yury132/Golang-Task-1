package http

import (
	"net/http"

	"github.com/Yury132/Golang-Task-1/internal/transport/http/handlers"
	"github.com/gorilla/mux"
)

func InitRoutes(h *handlers.Handler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", h.Home).Methods(http.MethodGet)
	r.HandleFunc("/auth", h.Auth).Methods(http.MethodGet)
	r.HandleFunc("/callback", h.Callback).Methods(http.MethodGet)
	r.HandleFunc("/me", h.Me).Methods(http.MethodGet)
	r.HandleFunc("/logout", h.Logout).Methods(http.MethodGet)
	r.HandleFunc("/users-list", h.GetUsersList).Methods(http.MethodGet)

	return r
}
