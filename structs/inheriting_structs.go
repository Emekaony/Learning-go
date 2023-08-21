package structs

import "fmt"

type Rectangle struct {
	width  int32
	height int32
}

type Square struct {
	width  int32
	height int32
}

type Shape struct {
	length int32
	width  int32
	name   string
	// structs can contain structs too!
	Rectangle
	Square
}

func (o *Rectangle) Area() int32 {
	return o.width * o.height
}

func (o *Square) Area() int32 {
	return o.height * o.width
}

func (o *Rectangle) Display() {
	fmt.Println(fmt.Sprintf("Rectangle: {width: %d, height: %d}", o.width, o.height))
}

func (o *Square) Display() {
	fmt.Println(fmt.Sprintf("Square: {width: %d, height: %d}", o.width, o.height))
}
