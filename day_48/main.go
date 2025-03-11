package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		select {
		case <-time.After(5 * time.Second):
			fmt.Println("Operacao concluida!")
		case <-ctx.Done():
			fmt.Println("Operacao cancelada!")
		}
	}(ctx)

	time.Sleep(2 * time.Second)
	cancel()

	time.Sleep(3 * time.Second)
}
