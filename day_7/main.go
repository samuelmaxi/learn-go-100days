package main

import "fmt"

func main() {

	for i := 1; i <= 10; i++ {
		fmt.Print("Tabuada do ", i, "\n")
		for j := 1; j <= 10; j++ {
			fmt.Println(i, "x", j, "=", i*j)
		}
		fmt.Print("==============\n")
	}
}
