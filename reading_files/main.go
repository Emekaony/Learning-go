package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	for _, a := range os.Args[1:] {
		fmt.Printf("--- %s ---\n", a)
		f, err := os.Open(a)
		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
			continue
		}
		io.Copy(os.Stdout, f)
	}
}
