package ranodmness

import (
	"fmt"
	"math/rand"
)

func Bla() {
	n := rand.Intn(100)
	fmt.Println(n)

	// shuffle a slice
	x := []string{"emeka", "chidera", "kamsi"}
	rand.Shuffle(len(x), func(i int, j int) {
		x[i], x[j] = x[j], x[i]
	})
	fmt.Println(x)
}

func main() {
	// generates a random number between 0 and 99
}
