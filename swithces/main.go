package main

import "fmt"

func main() {
	op := "windows"

	switch op {
	case "windows":
		fmt.Println("It is a windows op")
	case "ios":
		fmt.Println("ios op")
	default:
		fmt.Println("not sure what it is")
	}

}
