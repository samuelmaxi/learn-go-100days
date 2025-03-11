package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"strings"
)

type User struct {
	Name     string `json:"name"`
	PassWord string `json:"password"`
}

func main() {

	http.HandleFunc("/login", handlerLogin)

	slog.Info("Rodando na porta :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handlerLogin(w http.ResponseWriter, r *http.Request) {
	var u User

	auth := r.Header.Get("Authorization")

	fmt.Print(auth)
	if auth == "" || !strings.HasPrefix(auth, "Basic ") {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"error":"nao autorizado"}`))
		return
	}

	payload, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(auth, "Basic "))
	if err != nil {
		log.Fatal(err)
	}

	pair := strings.SplitN(string(payload), ":", 2)
	if len(pair) != 2 || pair[0] != "admin" || pair[1] != "123" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"error":"nao autorizado"}`))
		return
	}

	u.Name = pair[0]

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(u)
}
