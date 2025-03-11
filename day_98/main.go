package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {

	user := User{
		ID:   1,
		Name: "samuel",
	}

	byteValueJSON, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("user.json", byteValueJSON, 0644)
	if err != nil {
		fmt.Println(err)
	}

}
