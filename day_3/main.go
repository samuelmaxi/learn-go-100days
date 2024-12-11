package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	values := os.Args[1:]
	firstValue, _ := strconv.Atoi(values[0])
	secondValue, _ := strconv.Atoi(values[2])
	simbolValue := values[1]
	var result int

	switch {
	case simbolValue == "+":
		result = firstValue + secondValue
	case simbolValue == "-":
		result = firstValue - secondValue
	case strings.ToLower(simbolValue) == "x":
		result = firstValue * secondValue
	case simbolValue == "/":
		result = firstValue / secondValue
	}

	fmt.Println("valor: ", result)
}
