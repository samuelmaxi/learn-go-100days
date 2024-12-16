package main

import "fmt"

func main() {

	for i := range 11 {

		par := func() int {
			soma := i * 2
			return soma
		}()

		fmt.Println(par)
	}

}
