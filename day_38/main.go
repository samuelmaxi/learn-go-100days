package main

import (
	"log"
	"log/slog"
	"net/http"
)

func main() {
	// mux := http.NewServeMux()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://gole.com")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Content-Type", "application/json")

		w.Write([]byte("{\"hello\":\"world\"}"))
	})

	slog.Info("Rodando na porta :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
