package main

import "fmt"

type Talker interface {
	talk() string
}

type Cat struct{}

func (c *Cat) talk() string {
	return "Meouwwwww"
}

func do_talk(c Talker) string {
	// %#v means display the type of the object
	return fmt.Sprintf("The %#v says %s", c, c.talk())
}

func main() {
	fmt.Println(do_talk(&Cat{}))
}
