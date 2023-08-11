package main

import "fmt"

func main() {
	// basic way to make a map
	emap := make(map[string]int)
	fmt.Printf("%v\n", emap)

	// another way to make a map
	m := map[string]float32{"age": 22, "height": 6}
	m["bla"] = 99
	fmt.Println(m)

	for _, elem := range []string{"age", "height", "salinha", "bla"} {
		at, ok := m[elem]
		isThere := "exist"
		if !ok {
			isThere = "does not exist"
		}
		fmt.Printf("%s does %s in m at %f\n", elem, isThere, at)
	}
}
