package main

import "fmt"

func main() {
	sub := func(n int) int {
		return n - 7
	}
	fmt.Println("sub(10) is ", sub(10))
}

// closures are cool. Allowing you throw functions around
func AddNer(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}
