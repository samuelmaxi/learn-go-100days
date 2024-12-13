package main

import "fmt"

type studant struct {
	Nome  string
	Idade int
	Notas [4]float32
}

func main() {
	studant := studant{
		Nome:  "Samuel",
		Idade: 21,
		Notas: [4]float32{10.0, 7, 7, 9},
	}

	fmt.Println(studant)
}
