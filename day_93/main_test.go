package main

import "testing"

func TestCalcularDescontoEFormatarResultado(t *testing.T) {
	precoOriginal := 100.0
	percentualDesconto := 20.0

	precoComDesconto := CalcularDescont(precoOriginal, percentualDesconto)
	resultadoFormatado := Formatarresultado(precoOriginal, precoComDesconto)

	expected := "Preco original: 100.00, Preco com desconto: 80.00"

	if resultadoFormatado != expected {
		t.Errorf("Resultado esperado: %s, mas obteve: %s", expected, resultadoFormatado)
	}
}
