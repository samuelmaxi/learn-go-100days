package main

import "testing"

func TestAdd(t *testing.T) {
	result := Add(1, 2)
	if result != 3 {
		t.Errorf("Expected result to be 3, but got %d ", result)
	}
}
func TestAdd2(t *testing.T) {
	result := Add(2, 2)
	if result != 4 {
		t.Errorf("Expected result to be 4, but got %d ", result)
	}
}
func TestAdd3(t *testing.T) {
	result := Add(3, 2)
	if result != 5 {
		t.Errorf("Expected result to be 6, but got %d ", result)
	}
}
func TestAdd4(t *testing.T) {
	result := Add(21, 2)
	if result != 23 {
		t.Errorf("Expected result to be 23, but got %d ", result)
	}
}
