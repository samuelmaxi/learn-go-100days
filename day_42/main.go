package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {

	var wg sync.WaitGroup
	counter := Counter{}

	numGoroutunes := 100

	wg.Add(numGoroutunes)
	for i := 0; i < numGoroutunes; i++ {
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait()
	fmt.Printf("Valor final do contador: %d\n", counter.Value())
}
