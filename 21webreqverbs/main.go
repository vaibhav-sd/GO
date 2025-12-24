package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Web verb in GO")
	performGetRequest()
	performPostJsonRequest()
	performPostFormRequest()
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

func performPostJsonRequest() {
	const myurl = "http://localhost:8000/post"

	//json payload

	requestBody := strings.NewReader(`
		{
			"coursename" : "Let learn GO",
			"price":0
		}
	`)

	res, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	content, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println("Content : ", string(content))
}

func performPostFormRequest() {
	const myurl = "http://localhost:8000/postform"

	// formdata
	data := url.Values{}
	data.Add("firstname", "Vaibhav")
	data.Add("lastname", "dagwal")
	data.Add("email", "vaibhav@dagwal.com")

	res, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	content, _ := io.ReadAll(res.Body)

	fmt.Println(string(content))

}
