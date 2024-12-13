package main

import (
	"fmt"
	"math"
)

type Student struct {
	Nome  string
	Idade int
	Notas [4]float32
}

func (s Student) Media() float32 {
	var soma float32
	for i := 0; i < len(s.Notas); i++ {
		soma += s.Notas[i]
	}
	result := soma / float32(len(s.Notas))

	truncate := float32(math.Floor(float64(result*100)) / 100)

	return truncate
}

func main() {

	alunoA := Student{
		Nome:  "Samuel",
		Idade: 21,
		Notas: [4]float32{10.0, 7.6, 5.6, 9.9},
	}

	fmt.Println(alunoA.Media())
}
