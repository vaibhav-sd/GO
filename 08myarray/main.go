package main

import "fmt"

func main() {
	fmt.Println("Array in GO")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[2] = "Kiwi"

	fmt.Println("Fruit list is : ", fruitList)
	fmt.Println("Fruit list length is : ", len(fruitList))

	var vegList = [3]string{"potato", "beans", "mushroom"}

	fmt.Println("Veg list is : ", vegList)
	fmt.Println("Veg list length is : ", len(vegList))

}
