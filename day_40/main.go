package main

import (
	"fmt"
	"time"
)

func hello(s string) {
	for i := 0; i < 1; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go hello("Word!")
	go hello("Hello")
	time.Sleep(500 * time.Millisecond)
}
