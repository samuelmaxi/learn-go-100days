package main

import "fmt"

func Formatarresultado(precoOriginal, precoComDesconto float64) string {
	return fmt.Sprintf("Preco original: %.2f, Preco com desconto: %.2f", precoOriginal, precoComDesconto)
}
