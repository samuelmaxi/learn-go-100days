package main

import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 1
	}

	fmt.Println("Numero: ", n)

	return n * factorial(n-1)
}

func main() {
	fmt.Println(factorial(8))
}
