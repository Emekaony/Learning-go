package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	println("line by line")

	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		tx := in.Text()
		fmt.Printf("txt is %s\n", tx)
		if strings.Trim(tx, "\n ") == "next" {
			break
		}
	}
	println("goodbye")
}
