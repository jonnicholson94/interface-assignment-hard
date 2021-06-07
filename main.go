package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f := os.Args[1]

	readFile(f)
}

func readFile(l string) {
	f, err := os.Open(l)

	if err != nil {
		fmt.Println("There was an error reading the file:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)
}
