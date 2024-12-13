package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fmt.Println("Defer")
	if err := writer("readme.txt", "Ola.com.br"); err != nil {
		log.Fatal("Failed to write file: ", err)
	}

	fmt.Println("Panic")
	names := []string{"Samuel", "Maximo"}

	fmt.Println("Names: ", names[len(names)])
}

func writer(fileName, text string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.WriteString(file, text)
	if err != nil {
		return err
	}

	return err
}
