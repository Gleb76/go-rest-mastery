package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"ch02/internal/models"
)

func Health(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	resp := models.HealthResponse{
		Status:    "ok",
		Timestamp: time.Now().Format(time.RFC3339),
	}

	jsonData, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
