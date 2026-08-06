package api

import (
    "encoding/json"
    "log"
    "net/http"

    "tvremote/internal/model"
)

func writeJSON(w http.ResponseWriter, status int, resp model.APIResponse) {

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)

    if err := json.NewEncoder(w).Encode(resp); err != nil {
        log.Printf("encode response failed: %v", err)
    }
}