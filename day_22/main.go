package main

import "os"

func main() {
	err := os.Remove("doc.txt")
	if err != nil {
		panic(err)
	}
}
