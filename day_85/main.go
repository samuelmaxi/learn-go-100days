package main

import (
	"fmt"
	"time"
)

func myTask() {
	fmt.Println("Executando tarefa...", time.Now().Format("15:04:05"))
}

func main() {

	interval := 3 * time.Second

	for {
		go myTask()
		time.Sleep(interval)
	}

}
