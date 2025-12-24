package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Web verb in GO")
	performGetRequest()
}

func performGetRequest() {
	const myurl = "http://localhost:8000/get"

	res, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	fmt.Println("Status code: ", res.Status)
	fmt.Println("Content length is: ", res.ContentLength)

	var resString strings.Builder
	content, _ := io.ReadAll(res.Body)
	byteCount, _ := resString.Write(content)

	fmt.Println("ByteCount is: ", byteCount)
	fmt.Println(resString.String())

	// fmt.Println(string(content))

}
