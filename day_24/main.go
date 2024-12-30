package main

import (
	"fmt"
	"os"
)

type User struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func main() {
	f, err := os.ReadFile("user.json")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(f))
}
