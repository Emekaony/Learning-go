package main

import (
	"fmt"
)

func main() {
	var arr [5]int
	for a := 0; a < len(arr); a++ {
		arr[a] = a
	}
	var index int = 0
	for {
		if index == len(arr) {
			// say break when u want to leave the loop but not return!
			break
		}
		fmt.Printf("arr[%d] == %d", index, arr[index])
	}
}
