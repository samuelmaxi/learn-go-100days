package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	var u User

	f, err := os.ReadFile("user.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(f, &u)
	if err != nil {
		panic(err)
	}

	fmt.Print(u)

}
