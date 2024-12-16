package main

import "fmt"

func calc(list ...float32) float32 {
	var soma float32
	for i := 0; i < len(list); i++ {
		soma += list[i]
	}

	return soma
}

func main() {
	fmt.Println(calc(4, 1, 5, 76, 7.6))
}
