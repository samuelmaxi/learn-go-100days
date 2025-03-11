package main

import (
	"fmt"
	"strings"
)

func removeCharacter(word, character string) string {
	return strings.Trim(word, character)
}

func counterLetters(word, letter string) int {
	return strings.Count(word, letter)
}

func main() {

	fmt.Println(removeCharacter("??ABACAXI!!", "??!!"))
	fmt.Println(counterLetters("ABACAXI", "A"))

}
