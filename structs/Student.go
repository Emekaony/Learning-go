package structs

import "fmt"

// hide the struct, but make the make func visible
type Student struct {
	firstname  string
	lastname   string
	studentID  int16
	bestFriend string
}

func MakeNewStudent(fname string, lname string, id int16, bfriend string) *Student {
	// because we make a new struct in here
	// we must return a pointer to it so it doesn't die after the lifetime of this func
	s := Student{
		firstname:  fname,
		lastname:   lname,
		studentID:  id,
		bestFriend: bfriend,
	}
	fmt.Println("New student has been initialized.")
	return &s
}
