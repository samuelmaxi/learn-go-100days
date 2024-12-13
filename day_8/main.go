package main

import (
	"fmt"
	"reflect"
)

func main() {
	listFruits := []string{"apple", "mango", "watermelon", "cherry", "blue berries"}
	arrayFruits := make([]string, 5, 10)

	// ADD AND REMOVE
	listFruits = append(listFruits, "pineapple")
	listFruits = listFruits[:len(listFruits)-1]

	// ADD AND REMOVE
	arrayFruits = append(arrayFruits, "banana")
	arrayFruits = arrayFruits[:len(arrayFruits)-1]

	fmt.Println(reflect.TypeOf(listFruits))
	fmt.Println(listFruits)
	fmt.Println(reflect.TypeOf(arrayFruits))
	fmt.Println(arrayFruits)
}
