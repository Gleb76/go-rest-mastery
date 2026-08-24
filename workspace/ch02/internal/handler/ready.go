package handler

import (
	"encoding/json"
	"net/http"
	"sync"

	"ch02/internal/models"
)

var (
	ready    bool
	readyMux sync.RWMutex
)

func init() {
	ready = true
}

func Ready(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	readyMux.RLock()
	isReady := ready
	readyMux.RUnlock()

	if !isReady {
		http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
		return
	}

	resp := models.Response{
		Message: "ready",
		Status:  http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
