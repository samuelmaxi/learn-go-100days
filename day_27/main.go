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
	u := User{
		Name: "Samuel",
		Age:  21,
	}

	jsonData, err := json.Marshal(u)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("user.json", []byte(jsonData), 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println(jsonData)

}
