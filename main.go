package main

import (
	"log"
	"net/http"
	"time"
)

func handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte("{\"message\":\"Hello SLIIT!\"}\n"))
	})
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte("{\"status\":\"ok!!\"}\n"))
	})
	return mux
}

func main() {
	server := &http.Server{
		Addr:              ":8080",
		Handler:           handler(),
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
	}
	log.Println("Listening on :8080")
	log.Fatal(server.ListenAndServe())
}
