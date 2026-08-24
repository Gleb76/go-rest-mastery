package handler

import (
	"encoding/json"
	"net/http"
	"runtime"
	"time"

	"ch02/internal/models"
)

var startTime = time.Now()

func Info(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	uptime := time.Since(startTime).Seconds()
	resp := models.InfoResponse{
		Version:       "1.0.0",
		GoVersion:     runtime.Version(),
		UptimeSeconds: uptime,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
