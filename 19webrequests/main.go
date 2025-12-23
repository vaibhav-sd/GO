package main

import (
	"fmt"
	"io"
	"net/http"
)

const URL = "https://httpbin.org"

func main() {
	fmt.Println("Web requests in GO")
	response, err := http.Get(URL)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type: %T\n", response)
	defer response.Body.Close() // Closing the connection

	databytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	content := string(databytes)

	fmt.Println(content)
}
