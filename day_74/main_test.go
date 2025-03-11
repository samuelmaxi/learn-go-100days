package main

import "testing"

func TestSearchList(t *testing.T) {
	list := [2]string{"maca", "banana"}
	_, status := SearchList("banana", list[:])
	if !status {
		t.Errorf("Item nao encontrado na lista")
	}

}
