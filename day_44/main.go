package main

import "fmt"

func main() {
	ch := make(chan string, 2)

	ch <- "Golang"
	ch <- "Bufferred Channel"

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
