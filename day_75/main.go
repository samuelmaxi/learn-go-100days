package main

import (
	"fmt"
	"time"
)

func say(s string, done chan bool) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}

	done <- true
}

func main() {
	done := make(chan bool)

	go say("word", done)
	go say("hello", done)

	for i := 0; i < 2; i++ {
		<-done
	}
}
