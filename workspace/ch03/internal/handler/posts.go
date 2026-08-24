package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "пост создан (заглушка)",
	})
}

func GetUserPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "получены посты пользователя (заглушка)",
	})
}

func GetUserPostId(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "postID")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Post with id:",
		"id":      id,
	})
}
