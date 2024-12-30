package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type User struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func main() {
	http.HandleFunc("/", handlerListUser)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handlerListUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	f, err := os.ReadFile("user.json")
	if err != nil {
		panic(err)
	}

	var usersData []User
	err = json.Unmarshal(f, &usersData)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(usersData)
}
