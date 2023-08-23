package main

import (
	"fmt"
	"time"
)

/*
time.Format() takes a time.Time and converts to string
time.Parse() does the direct opposite
*/
func main() {
	// currentTime := time.Now()
	// fmt.Println("The current time is: ", currentTime)
	// fmt.Println("We are currently in the month of", currentTime.Month())
	// fmt.Printf("The time is %d:%d\n", currentTime.Hour(), currentTime.Minute())

	// using format
	// theTime := time.Date(2023, 8, 15, 4, 21, 45, 100, time.Local)
	// another := time.Date(2001, 7, 25, 13, 30, 14, 12, time.Local)
	// fmt.Println("The time is", theTime)
	// fmt.Println(theTime.Format("2006-01-02 3:4:5 pm"))
	// fmt.Println(another.Format("2006-01-02 15:4:5 pm"))
	// // just a way to return date as a string using well-known RFC format
	// fmt.Println(another.Format(time.RFC3339Nano))

	// parsing date and time
	timeString := "7/25/2019 13:45:00"
	theTime, err := time.Parse("1/02/2006 15:04:05", timeString)
	if err != nil {
		fmt.Println("Could not parse time")
	}
	fmt.Println("The time is", theTime)
	fmt.Println(theTime.Format(time.RFC3339Nano))
}
