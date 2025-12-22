package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Files in GO")
	content := "This is need to go in file"

	file, err := os.Create("./myfile.txt")
	checkNilErr(err)

	length, err := io.WriteString(file, content)

	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)

	fmt.Println("Length is : ", length)
	defer file.Close()

	readFile("./myfile.txt")
}

func readFile(filename string) {
	data, err := os.ReadFile(filename)
	checkNilErr(err)

	fmt.Println("Text data in file is : ", data)
	fmt.Println("Text data in file is : ", string(data))

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
