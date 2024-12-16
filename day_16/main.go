package main

import "fmt"

func main() {
	a := "asdad"
	a = "asd"
	b := &a

	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(b)
	fmt.Println(*b)

}
