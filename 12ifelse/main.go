package main

import "fmt"

func main() {
	fmt.Println("If Else in GO")

	loginCount := 19
	var result string

	if loginCount < 10 {
		result = "Regular"
	} else {
		result = "Else"
	}

	fmt.Println(result)

	if num := 1; num < 10 {
		fmt.Println("Less than 10")
	} else {
		fmt.Println("Greater than 10")
	}
}
