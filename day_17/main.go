package main

import "fmt"

func fibonacci(i int) int {
	if i <= 1 {
		return i
	}
	return fibonacci(i-1) + fibonacci(i-2)
}

func main() {

	for i := 0; i < 10; i++ {

		fmt.Println("fibonacci", fibonacci(i))
	}

}
