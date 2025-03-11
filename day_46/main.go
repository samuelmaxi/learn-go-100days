package main

import (
	"fmt"
	"sync"
	"time"
)

func check(u string, checked chan<- bool) {
	time.Sleep(4 * time.Second)
	checked <- true
}

func IsReachable(urls []string) bool {
	ch := make(chan bool, len(urls))
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)

		go func(u string) {
			defer wg.Done()
			select {
			case <-time.After(1 * time.Second):
				ch <- false
			case <-func() <-chan bool {
				done := make(chan bool, 1)
				go func() {
					check(u, done)
				}()
				return done
			}():
				ch <- true
			}
		}(url)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for result := range ch {
		if result {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(IsReachable([]string{"url1", "url2"}))
}
