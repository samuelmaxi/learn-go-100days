package main

import (
	"fmt"
	"math"
)

// REFATORACAO DO CODIGO DO DIA 13

type Student struct {
	Nome  string
	Idade int
	Notas []float32
}

func (s Student) Media() float32 {
	var soma float32
	for _, nota := range s.Notas {
		soma += nota
	}
	result := soma / float32(len(s.Notas))

	return float32(math.Floor(float64(result*100)) / 100)
}

func main() {

	alunoA := Student{
		Nome:  "Samuel",
		Idade: 21,
		Notas: []float32{10.0, 7.6, 5.6, 9.9},
	}

	fmt.Println(alunoA.Media())
}
