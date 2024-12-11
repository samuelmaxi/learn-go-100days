package main

import (
	"fmt"
)

type User struct {
	Nome  string
	Idade int
}

func main() {
	numeroInteiro := 120
	numeroFlutuante := 3.1412
	palavra := "Abacaxi"
	lista := []any{1, 2, 3, "um", "dois", "tres"}
	pessoa := User{
		Nome:  "Samuel",
		Idade: 21,
	}
	marcado := false

	fmt.Println(numeroInteiro)
	fmt.Println(numeroFlutuante)
	fmt.Println(palavra)
	fmt.Println(lista)
	fmt.Println(pessoa)
	fmt.Println(marcado)
}
