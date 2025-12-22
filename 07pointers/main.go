package main

import "fmt"

func main() {
	fmt.Println("Pointers in GO")

	// var ptr *int
	// fmt.Println("Value of pointer is ", ptr)

	myNumber := 23

	var ptr = &myNumber
	fmt.Println("Value ot actual pointer is", ptr)
	fmt.Println("Value ot actual pointer is", *ptr)

	*ptr = *ptr + 2

	fmt.Println("New value is ", *ptr)
}
