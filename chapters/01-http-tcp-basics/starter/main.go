// Starter — Глава 01: HTTP основы
// Запуск: go run ./chapters/01-http-tcp-basics/starter/main.go
// Или: make run-ch01
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprint(w, "Go REST Mastery — Chapter 01\n")
	})

	mux.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, HTTP!\n")
	})

	mux.HandleFunc("GET /info", func(w http.ResponseWriter, r *http.Request) {
		info := map[string]string{
			"method": r.Method,
			"path":   r.URL.Path,
			"host":   r.Host,
			"proto":  r.Proto,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(info)
	})

	mux.HandleFunc("POST /echo", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		buf := make([]byte, r.ContentLength)
		if r.ContentLength > 0 {
			r.Body.Read(buf)
		}
		w.Write(buf)
	})

	addr := ":8080"
	log.Printf("Chapter 01 starter listening on http://localhost%s", addr)
	log.Printf("Try: curl -v http://localhost%s/info", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
