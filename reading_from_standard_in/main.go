package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/crypto/ssh/terminal"
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
	println("RawMode")
	raw()
	println("goodbye")

}

func raw() {
	oldstate, err := terminal.MakeRaw(0)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer terminal.Restore(0, oldstate)
	reader := bufio.NewReader(os.Stdin)

	for {
		c, _, err := reader.ReadRune()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Rune = %c\n\r", c)
		if c == 'q' {
			return
		}
	}
}
