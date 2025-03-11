package main

import (
	"fmt"
	"os"
)

func main() {
	url := os.Getenv("URL")

	fmt.Println("url: ", url)
}
