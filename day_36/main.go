package main

import (
	"encoding/json"
	"io"
	"log"
	"log/slog"
	"net/http"
)

type User struct {
	Name string `json:"name"`
}

func main() {
	http.HandleFunc("/", handlerPost)
	slog.Info("Rodando na porta :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handlerPost(w http.ResponseWriter, r *http.Request) {
	var u User

	jsonData, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	err = json.Unmarshal(jsonData, &u)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(u)
}
