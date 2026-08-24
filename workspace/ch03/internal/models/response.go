package models

type Response struct {
	Message    string
	StatusCode int
}

type Article struct {
	Title string `json:"title"`
	ID    string `json:"id"`
}
