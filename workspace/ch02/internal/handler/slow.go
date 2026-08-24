package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"ch02/internal/models"
)

func Slow(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	secondsStr := r.URL.Query().Get("seconds")
	seconds, err := strconv.Atoi(secondsStr)
	if err != nil || seconds <= 0 {
		seconds = 5
	}

	done := make(chan bool)
	go func() {
		time.Sleep(time.Duration(seconds) * time.Second)
		done <- true
	}()

	select {
	case <-done:
		resp := models.Response{
			Message: fmt.Sprintf("Slept for %d seconds", seconds),
			Status:  http.StatusOK,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	case <-r.Context().Done():
		http.Error(w, "Request timeout", http.StatusGatewayTimeout)
	}
}
