package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Book struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	fmt.Println("Rodando...")

	http.HandleFunc("/", handlerGetCat200)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handlerGetCat200(w http.ResponseWriter, r *http.Request) {
	req, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		panic(err)
	}

	// client := &http.Client{}

	// resp, err := client.Do()
	// if err != nil {
	// 	panic(err)
	// }

	// defer resp.Body.Close()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	var b []Book
	err = json.Unmarshal(body, &b)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
	w.Header()
}
