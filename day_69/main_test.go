package main

import "testing"

func TestAddInt(t *testing.T) {
	result := Add(1, 2)
	if result != 3 {
		t.Fatal("Deve ser 3")
	}
}
