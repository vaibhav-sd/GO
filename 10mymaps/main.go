package main

import "fmt"

func main() {
	fmt.Println("Maps in GO")

	langs := make(map[string]string)

	langs["JS"] = "JavaScript"
	langs["RB"] = "Ruby"
	langs["PY"] = "Python"

	fmt.Println(langs)
	fmt.Println("JS shorts for :", langs["JS"])

	delete(langs, "RB")
	fmt.Println(langs)

	// Loops

	for key, value := range langs {
		fmt.Printf("For key %v value is %v\n", key, value)
	}
}
