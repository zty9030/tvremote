package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"tvremote/internal/api"
)

func RegisterRoutes(r *chi.Mux, h *api.Handler) {
	r.Post("/api/key", h.Key)

	fs := http.FileServer(http.Dir("./web"))

	r.Handle("/*", fs)

}
