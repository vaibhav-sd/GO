package main

import (
	"fmt"
	"net/url"
)

const myurl string = "http://www.example.com/path?id=1&name=test"

func main() {
	fmt.Println("Handling urls in GO")

	result, err := url.Parse(myurl)

	if err != nil {
		panic(err)
	}

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	// fmt.Println(result.RawQuery)

	qparams := result.Query()

	fmt.Printf("The type of query params are: %T\n", qparams)

	fmt.Println(qparams["name"])

	for _, val := range qparams {
		fmt.Println("Param is: ", val)
	}

	partsOfUrl := &url.URL{
		Scheme:   "http",
		Host:     "example.com",
		Path:     "/path",
		RawQuery: "name=test",
	}

	params := partsOfUrl.Query()
	params.Add("id", "1")
	partsOfUrl.RawQuery = params.Encode()

	anotherURL := partsOfUrl.String()

	fmt.Println(anotherURL)

}
