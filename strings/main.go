package strings

import (
	"fmt"
	"strings"
)

func main() {
	s := "nnaemeka onyeokoro"
	for k, v := range s {
		fmt.Printf("[%d] = %c\n", k, v)
	}

	// join strings together using slices
	j := strings.Join([]string{"orange", "pineapple", "watermellon"}, ", ")
	fmt.Println(j)

	// split a string into an array
	k := strings.Split(j, ", ")
	fmt.Println(k)
}
