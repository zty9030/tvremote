package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func registerRoutes(r *chi.Mux) {

	fs := http.FileServer(http.Dir("./web"))

	r.Handle("/*", fs)

}