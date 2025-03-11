package main

import "testing"

func TestSoma(t *testing.T) {
	result := Soma(1, 4)
	if result != 3 {
		t.Error("Nao deu 3")
	}
}
