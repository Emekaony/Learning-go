package main

import (
	"fmt"
	"fundamentals/structs"
)

func main() {

	fname := "emeka"
	lname := "onyeokoro"
	id := 1234
	bfriend := "chizalu"

	fmt.Println(structs.MakeNewStudent(fname, lname, int16(id), bfriend))
}
