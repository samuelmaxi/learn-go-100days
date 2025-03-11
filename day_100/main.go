package main

import (
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("./texto.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))

}
