package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog"
)

type Handler struct {
	log zerolog.Logger
}

func New(log zerolog.Logger) *Handler {
	return &Handler{
		log: log,
	}
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data := "{\"health\": \"ok\"}"

	response, err := json.Marshal(data)
	if err != nil {
		h.log.Error().Err(err).Msg("filed to marshal response data")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(response)
}

func (h *Handler) Hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte("hello"))
}
