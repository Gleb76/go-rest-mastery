package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
)

type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type HeadersResponse struct {
	Headers map[string]string `json:"headers"`
}

type QueryResponse struct {
	Query string `json:"query"`
	Page  string `json:"page"`
}

type ErrorResponse struct {
	Error string `json:"error"`
	Path  string `json:"path"`
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", handleRoot)
	mux.HandleFunc("/hello", handleHello)
	mux.HandleFunc("POST /echo", handleEchoBody)
	mux.HandleFunc("ANY /method", handleAnyMethod)
	mux.HandleFunc("GET /headers", handleHeaders)
	mux.HandleFunc("GET /status/", handleStatus)
	mux.HandleFunc("GET /search", handleSearch)
	mux.HandleFunc("/{any}", handleNotFound)

	log.Fatal(http.ListenAndServe(":8080", mux))
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	wc, err := w.Write([]byte("Go REST Mastery"))
	if err != nil {
		slog.Error("error writing response", "err", err)
		return
	}
	fmt.Printf("%d bytes written\n", wc)
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	wc, err := w.Write([]byte("Hello, HTTP!"))
	if err != nil {
		slog.Error("error writing response", "err", err)
		return

	}
	fmt.Printf("%d bytes written\n", wc)
}

func handleEchoBody(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		slog.Error("error reading body", "err", err)
		http.Error(w, "Failed to read body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	wc, err := w.Write(body)
	if err != nil {
		slog.Error("error writing response", "err", err)
		return

	}
	fmt.Printf("%d bytes written\n", wc)
}

func handleAnyMethod(w http.ResponseWriter, r *http.Request) {
	data := Response{
		Message: "Hello, JSON!",
		Status:  200,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	wc, err := w.Write(jsonData)
	if err != nil {
		slog.Error("error writing response", "err", err)
		return
	}
	fmt.Printf("%d bytes written\n", wc)
}

func handleHeaders(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	headers := make(map[string]string)
	for name, values := range r.Header {
		if len(values) == 1 {
			headers[name] = values[0]
		} else {
			for i, value := range values {
				if i == 0 {
					headers[name] = value
				} else {
					headers[name] += ", " + value
				}
			}
		}
	}

	response := HeadersResponse{
		Headers: headers,
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		slog.Error("error marshaling JSON", "err", err)
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	wc, err := w.Write(jsonData)
	if err != nil {
		slog.Error("error writing response", "err", err)
		return
	}
	fmt.Printf("%d bytes written\n", wc)
}

func handleStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	path := strings.TrimPrefix(r.URL.Path, "/status/")

	statusCode, err := strconv.Atoi(path)
	if err != nil {
		http.Error(w, "Invalid status code", http.StatusBadRequest)
		return
	}

	if statusCode < 100 || statusCode > 599 {
		http.Error(w, "Status code must be between 100 and 599", http.StatusBadRequest)
		return
	}

	statusText := http.StatusText(statusCode)
	if statusText == "" {
		statusText = "Unknown Status"
	}

	response := Response{
		Message: fmt.Sprintf("Status %d - %s", statusCode, statusText),
		Status:  statusCode,
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		slog.Error("error marshaling JSON", "err", err)
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	wc, err := w.Write(jsonData)
	if err != nil {
		slog.Error("error writing response", "err", err)
		return
	}
	fmt.Printf("%d bytes written with status %d\n", wc, statusCode)
}

func handleSearch(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	queryParams := r.URL.Query()

	query := queryParams.Get("q")
	page := queryParams.Get("page")

	response := QueryResponse{
		Query: query,
		Page:  page,
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		slog.Error("error marshaling JSON", "err", err)
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	wc, err := w.Write(jsonData)
	if err != nil {
		slog.Error("error writing response", "err", err)
		return
	}
	fmt.Printf("%d bytes written\n", wc)
}

func handleNotFound(w http.ResponseWriter, r *http.Request) {
	errResponse := ErrorResponse{
		Error: "not found",
		Path:  r.URL.Path,
	}

	jsonData, err := json.Marshal(errResponse)
	if err != nil {
		slog.Error("error marshaling JSON", "err", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)

	wc, err := w.Write(jsonData)
	if err != nil {
		slog.Error("error writing response", "err", err)
		return
	}
	fmt.Printf("%d bytes written (404 not found: %s)\n", wc, r.URL.Path)
}

// ## Бонус ⭐⭐⭐⭐

// TCP-сервер без HTTP (чистый `net`):
// - Слушай `:9000`
// - При подключении отправь `220 Welcome\r\n`
// - Прочитай одну строку от клиента и отправь `250 OK: <строка>\r\n`

// ```bash
// nc localhost 9000
// ```
