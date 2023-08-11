package main

import (
	"fmt"
)

func main() {
	a := [10]int32{1, 2, 3}
	view := a[2:]

	// because view is a slice of a, changing it will
	// also alter the parent array, a
	view[3] = 100
	// fmt.Println(view)
	// fmt.Println(a)

	// but when we make view bigger, it will have to get allocated
	// more memory space, hence dicsontinue to point to the same mem loc as a
	// so now changing it after giving it more space will have no effect on a
	view = append(view, 32)
	fmt.Println(view)
	fmt.Println(a) // a will not have 32 because view is pointing to somewhere new in mem

	// append a slice into another slice
	parent := []float32{1, 2, 3}
	child := []float32{4, 5, 6}
	fmt.Println(parent)
	fmt.Println(child)

	// now append child to parent using the spread operator
	parent = append(parent, child...)
	fmt.Println("Parent is now: ", parent)
	fmt.Println("child is still: ", child)
}
