package main

import (
	"fmt"
)

type Point struct {
	x int
	y int
}

// functions defined for structs use pointers
// for memory I guess. Will figure out later
func (p *Point) move(x int, y int) {
	p.x += x
	p.y += y
}

func (p *Point) display() string {
	/*
		Printf just prints to the output
		Sprintf saves it into a string so we can return it
	*/
	return fmt.Sprintf("{x: %d, y: %d}", p.x, p.y)
}

func main() {
	var t Point
	p := Point{x: 0, y: 0}
	p.move(3, 3)
	fmt.Println(p.display())
	t.x = 10
	t.y = 24
	// display returns a string so we save it in a variable
	var k string = t.display()
	fmt.Println(k)

	s1 :=
		Shape{width: 10, length: 44,
			Rectangle: Rectangle{width: 55, height: 12},
			Square:    Square{width: 19, height: 99},
		}
	// go was intelligently built
	/*
		when u have two structs inside a struct with methods of the same name
		you must call the specific struct first before calling the method they both share in common
		this helps avoid confusion for the compiler
	*/
	s1.Rectangle.Area()
	s1.Square.Area()
	s1.Rectangle.Display()
	s1.Square.Display()

	car1 := Car{color: "red", Benz: Benz{name: "GLE 450", year: 2023, owner: "Nnaemeka Onyeokoro"}}
	car1.display()
}
