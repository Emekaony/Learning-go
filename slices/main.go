package main

import "fmt"

func main() {
	var tt = []int{1, 2, 3, 4}
	var ll = tt[:]

	fmt.Printf("Before changing ll, tt[0] is %d\nAnd ll[0] is %d\n", tt[0], ll[0])
	ll[0] = 99
	fmt.Printf("After changing ll, tt[0] is %d\nAnd ll[0] is %d\n", tt[0], ll[0])

	ll = append(ll, 200) // points to a different location in memory
	fmt.Println("We have appened to ll, now it points to a diff loc in memory")
	ll[0] = 501

	fmt.Printf("After pointing ll to a new loc in memory and changing ll[0], it should not affect tt. ll[0] is now %d and tt[0] is still %d.\n", ll[0], tt[0])
}
