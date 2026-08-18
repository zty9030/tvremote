package api

import (
    "encoding/json"
    "net/http"

    "tvremote/internal/input"
    "tvremote/internal/model"
)

type Handler struct {
    input *input.Service
}

func NewHandler(input *input.Service) *Handler {
    return &Handler{
        input: input,
    }
}

// Key handles a normal key press.
// POST /api/key
func (h *Handler) Key(w http.ResponseWriter, r *http.Request) {
    var req model.KeyRequest

    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        writeJSON(w, http.StatusBadRequest, model.APIResponse{
            Success: false,
            Message: err.Error(),
        })
        return
    }

    if err := h.input.SendKey(req.Key); err != nil {
        writeJSON(w, http.StatusInternalServerError, model.APIResponse{
            Success: false,
            Message: err.Error(),
        })
        return
    }

    writeJSON(w, http.StatusOK, model.APIResponse{
        Success: true,
        Message: "OK",
    })
}

// KeyDown handles a key-down event.
// POST /api/key/down
func (h *Handler) KeyDown(w http.ResponseWriter, r *http.Request) {
    var req model.KeyRequest

    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        writeJSON(w, http.StatusBadRequest, model.APIResponse{
            Success: false,
            Message: err.Error(),
        })
        return
    }

    if err := h.input.KeyDown(req.Key); err != nil {
        writeJSON(w, http.StatusInternalServerError, model.APIResponse{
            Success: false,
            Message: err.Error(),
        })
        return
    }

    writeJSON(w, http.StatusOK, model.APIResponse{
        Success: true,
        Message: "OK",
    })
}

// KeyUp handles a key-up event.
// POST /api/key/up
func (h *Handler) KeyUp(w http.ResponseWriter, r *http.Request) {
    var req model.KeyRequest

    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        writeJSON(w, http.StatusBadRequest, model.APIResponse{
            Success: false,
            Message: err.Error(),
        })
        return
    }

    if err := h.input.KeyUp(req.Key); err != nil {
        writeJSON(w, http.StatusInternalServerError, model.APIResponse{
            Success: false,
            Message: err.Error(),
        })
        return
    }

    writeJSON(w, http.StatusOK, model.APIResponse{
        Success: true,
        Message: "OK",
    })
}
