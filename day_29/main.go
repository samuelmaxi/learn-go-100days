package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlerHelloWord)
	fmt.Println("Rodando...")

	log.Fatal(http.ListenAndServe(":8081", nil))
}

func handlerHelloWord(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"hello": "word"}`))
}
