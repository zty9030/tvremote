package api

import (
    "encoding/json"
    "log"
    "net/http"

    "tvremote/internal/model"
)

func Key(w http.ResponseWriter, r *http.Request) {

    var req model.KeyRequest

    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {

        http.Error(w, err.Error(), http.StatusBadRequest)

        return
    }

    log.Printf("KEY = %s\n", req.Key)
    
    resp := model.APIResponse{
        Success: true,
        Message: "OK",
    }

    w.Header().Set("Content-Type", "application/json")

    if err := json.NewEncoder(w).Encode(resp); err != nil {
	log.Printf("encode response failed: %v", err)
    }
}
