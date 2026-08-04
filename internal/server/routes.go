package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"tvremote/internal/api"
)

func registerRoutes(r *chi.Mux) {

	r.Post("/api/key", api.Key)

	fs := http.FileServer(http.Dir("./web"))

	r.Handle("/*", fs)

}
