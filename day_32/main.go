package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"time"
)

func loggerHandler(next http.Handler) http.Handler {
	return http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		page := fmt.Sprintf("Pagina no caminho: %s", r.URL.Path)

		user, _ := r.Cookie("user")
		if r.URL.Path == "/admin" && user.Value != "admin" {
			w.Write([]byte("NAO AUTORIZADO"))
			fmt.Println(page + " | NAO AUTORIZADO")
			return
		}
		fmt.Println(page)
		next.ServeHTTP(w, r)
	}))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", whatsPage)
	slog.Info("Servidor rodando na porta :8080")

	logger := loggerHandler(mux)

	log.Fatal(http.ListenAndServe(":8080", logger))
}

func whatsPage(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{Name: "user", Value: "admin", Expires: time.Now().Add(1 * time.Minute), HttpOnly: true}
	http.SetCookie(w, &cookie)

	w.Write([]byte("HELLO WORD!"))
}
