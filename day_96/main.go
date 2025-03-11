package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	dateFlag := flag.Bool("date", false, "Exibir a data atual")
	flag.Parse()

	if *dateFlag {
		currentTime := time.Now()
		fmt.Println("Data atual", currentTime.Format("2006-01-02 15:04:05"))
	}
}
