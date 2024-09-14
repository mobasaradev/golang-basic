package main

import "fmt"

func main()  {
	// Slices was created with make function
	// make func has type, size
	// append func add new value to the list
	name := make([]string, 0)
	names := append(name, "lucky")
	fmt.Println("Names :", names)

	// Slices - length will be dynamic
	// bookingSeats := []string{} //value initialize
	var bookingSeats []string //later value initialize

	// Slices - no need to mention the index
	bookingSeats = append(bookingSeats, "Sara")
	fmt.Printf("The Students whom booked the seats are %v.\n", bookingSeats)
}