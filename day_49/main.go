package main

import (
	"fmt"
)

func createNumber(ch chan<- int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}

func filterPar(chIn <-chan int, ch chan<- int) {

	for i := range chIn {
		if i%2 == 0 {
			ch <- i
		}
	}
	close(ch)
}

func squareNumbers(chIn <-chan int, ch chan<- int) {
	for num := range chIn {
		ch <- num * num
	}
	close(ch)
}

func main() {
	numbers := make(chan int)
	eventNumbers := make(chan int)
	sqaredNumers := make(chan int)

	go createNumber(numbers)
	go filterPar(numbers, eventNumbers)
	go squareNumbers(eventNumbers, sqaredNumers)

	for result := range sqaredNumers {
		fmt.Println(result)
	}

}
