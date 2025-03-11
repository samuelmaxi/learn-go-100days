package main

import "fmt"

func SearchList(itemSearch string, listItens []string) (string, bool) {
	for i := 0; i < len(listItens); i++ {
		if listItens[i] == itemSearch {
			result := fmt.Sprintf("Item encontrado: %s", listItens[i])
			return result, true
		}
	}
	result := fmt.Sprintln("Item NAO encontrado")
	return result, false
}

func main() {
}
