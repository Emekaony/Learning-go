package structs

import "fmt"

type Car struct {
	color string
	Benz
}

type Benz struct {
	name  string
	year  int32
	owner string
}

func (car *Car) display() {
	// becuse Car has a struct of Benz that has a property called owner
	// it inherits that property from it. This is called struct embedding in go
	fmt.Printf("The owner of the car is %s, and the color is %s\n", car.owner, car.color)
}

func (b *Benz) display() {
	fmt.Println("The owner of the benz is ", b.owner)
}
