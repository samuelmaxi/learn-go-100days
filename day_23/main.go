package main

import (
	"encoding/json"
	"os"
)

type User struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func main() {

	users := []User{
		{ID: 1, Name: "Samuel"},
		{ID: 2, Name: "João"},
	}

	json, err := json.Marshal(users)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("user.json", json, 0660)
	if err != nil {
		panic(err)
	}

}
