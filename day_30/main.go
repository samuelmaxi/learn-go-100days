package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", handlerGET)
	http.HandleFunc("/post", handlerPOST)
	fmt.Println("Rodando...")

	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func handlerGET(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET | 200")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "get request!}`))
}

func handlerPOST(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST | 200")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "post request!}`))
}
