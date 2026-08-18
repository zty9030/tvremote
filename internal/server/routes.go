package server

import (
    "io/fs"
    "net/http"

    "github.com/go-chi/chi/v5"

    "tvremote/internal/api"
    webassets "tvremote/web"
)

func RegisterRoutes(r *chi.Mux, h *api.Handler) {

    r.Post("/api/key", h.Key)
    r.Post("/api/key/down", h.KeyDown)
    r.Post("/api/key/up", h.KeyUp)

    staticFS, err := fs.Sub(webassets.Assets, ".")
    if err != nil {
        panic(err)
    }

    r.Handle("/*", http.FileServer(http.FS(staticFS)))
}