package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"tvremote/internal/api"
)

func New(handler *api.Handler) http.Handler {

	r := chi.NewRouter()

	RegisterRoutes(r, handler)

	return r
}
