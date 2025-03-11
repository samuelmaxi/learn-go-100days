package main

import (
	"fmt"
	"time"
)

func process() {
	start := time.Now()
	time.Sleep(2 * time.Second)
	elapsed := time.Since(start)

	fmt.Println("Tempo de execução: ", elapsed)
}

func main() {

	process()

}
