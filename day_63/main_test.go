package main

import "testing"

func BenchmarkAddNames(b *testing.B) {
	names := []string{"Tiago", "Alexandre", "Luara"}
	people := []People{
		{Name: "Matheus", Age: 21},
		{Name: "Lucas", Age: 58},
		{Name: "Rafael", Age: 57},
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = AddNames(names, people)
	}
}
