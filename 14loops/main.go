package main

import "fmt"

func main() {
	fmt.Println("Loops in GO")

	days := []string{"Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	for ind, day := range days {
		fmt.Printf("Index is %v and day is %v \n", ind, day)
	}

	j := 1

	for j < 10 {
		if j == 5 {
			j++
			continue
		}
		fmt.Println(j)
		j++
	}
}
