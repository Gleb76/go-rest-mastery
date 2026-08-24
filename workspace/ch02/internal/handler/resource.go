package handler

import (
	"encoding/json"
	"net/http"

	"ch02/internal/models"
)

func Resource(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		resourceGet(w, r)
	case http.MethodPost:
		resourcePost(w, r)
	case http.MethodDelete:
		resourceDelete(w, r)
	default:
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
	}
}

func resourceGet(w http.ResponseWriter, r *http.Request) {
	resp := models.Response{
		Message: "Resource retrieved successfully",
		Status:  http.StatusOK,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func resourcePost(w http.ResponseWriter, r *http.Request) {
	resp := models.Response{
		Message: "Resource created successfully",
		Status:  http.StatusCreated,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

func resourceDelete(w http.ResponseWriter, r *http.Request) {
	resp := models.Response{
		Message: "Resource deleted successfully",
		Status:  http.StatusOK,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
