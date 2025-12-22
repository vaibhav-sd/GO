package main

import "fmt"

const LOGIN_TOKEN = "djkafldaffd"

func main() {
	var username string = "Vaibhav"
	fmt.Println(username)
	fmt.Printf("Variable type is : %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable type is : %T \n", isLoggedIn)

	var smallVal uint8 = 8
	fmt.Println(smallVal)
	fmt.Printf("Variable type is : %T \n", smallVal)

	var smallFloat float32 = 383.33
	fmt.Println(smallFloat)
	fmt.Printf("Variable type is : %T \n", smallFloat)

	var variable int
	fmt.Println(variable)
	fmt.Printf("Variable type is : %T \n", variable)

	fmt.Printf(LOGIN_TOKEN)

}
