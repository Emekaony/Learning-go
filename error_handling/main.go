package main

import "fmt"

func safeDivision(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("Cannot divide by Zero")
	}
	return a / b, nil
}

func main() {
	const num1 int = 12
	const den1 int = 0
	const den2 int = 4

	// call the function
	res1, err1 := safeDivision(num1, den1)
	res2, err2 := safeDivision(num1, den2)

	fmt.Printf("Dividing %d by %d yields %d\nThe error involved is `%s`", num1, den1, res1, err1)
	fmt.Printf("Dividing %d by %d yields %d\nThe error involved is `%s`", num1, den2, res2, err2)
}
