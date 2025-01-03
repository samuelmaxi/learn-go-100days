package main

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
)

func main() {

	http.HandleFunc("/", handlerGetNumber)
	http.HandleFunc("/query", handlerGetNumber)
	fmt.Println("Rodando...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handlerGetNumber(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	param1 := query.Get("p")

	path := fmt.Sprintf("{messageID:%s}", param1)

	fmt.Print(reflect.TypeOf(path))
	fmt.Print(path)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(path))
}
