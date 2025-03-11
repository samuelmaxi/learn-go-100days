package main

import (
	"fmt"
	"sync"
	"time"
)

func task(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Iniciando tarefa %d... \n", id)
	time.Sleep(time.Duration(id) * time.Second)
	fmt.Printf("Fim da tarefa %d... \n", id)

}

func main() {
	var wg sync.WaitGroup

	numTasks := 3
	wg.Add(numTasks)

	for i := 1; i <= numTasks; i++ {
		go task(i, &wg)
	}

	wg.Wait()
	fmt.Println("todas tarefas foram concluidas")
}
