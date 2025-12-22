package main

import "fmt"

func main() {
	fmt.Println("Functions is GO")
	greeter()

	// greeterTwo()

	result := adder(3, 9)

	fmt.Println("Result is ", result)
	proResult, proMsg := proAdder(2, 284, 83)

	fmt.Printf("Pro Adder %v, Message is %v ", proResult, proMsg)
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}

	return total, "Result from Pro Function"
}

// func greeterTwo() {
// 	fmt.Println("Hello from second function")
// }

func greeter() {
	fmt.Println("Hello")
}
