package main

import "fmt"

type Printer interface {
	Print() string
}

type GeometricShapes struct {
	Name string
}

type Circulo struct {
	GeometricShapes
}

type Retagule struct {
	GeometricShapes
}

func (g *GeometricShapes) Print() string {
	return fmt.Sprintf("Nome: %s", g.Name)
}

func display(p Printer) {
	fmt.Println(p)
	fmt.Println(p.Print())
}

func main() {
	circule := &Circulo{GeometricShapes{Name: "Circulo"}}
	retangule := &Retagule{GeometricShapes{Name: "Retangule"}}

	display(circule)
	display(retangule)
}
