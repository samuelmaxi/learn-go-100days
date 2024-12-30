package main

import "os"

func main() {

	args := os.Args[1]
	t := []byte(args)

	err := os.WriteFile("doc.txt", t, 0700)
	if err != nil {
		panic(err)
	}

}
