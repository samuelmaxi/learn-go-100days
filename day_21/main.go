package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.ReadFile("doc.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(file))
}
