package main

import (
	"fmt"
)

var i interface{} = "aaa"

func main() {
	t, ok := i.(string)

	fmt.Println(t, ok)
}
