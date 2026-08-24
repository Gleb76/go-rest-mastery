package handler

import (
	"encoding/json"
	"net/http"
)

func AdminReady(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	readyMux.Lock()
	ready = !ready
	newState := ready
	readyMux.Unlock()

	resp := map[string]bool{"ready": newState}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
