package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`              // - will not show the password
	Tags     []string `json:"tags,omitempty"` // omitempty will not show empty field
}

func main() {
	fmt.Println("JSON in GO")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	courses := []course{
		{"GO", 299, "yt", "abc123", []string{"web", "backend"}},
		{"Python", 299, "yt", "bcd123", []string{"web", "ML", "backend"}},
		{"Java", 399, "yt", "ijk123", nil},
	}

	// package this data into JSON

	// finalJson, err := json.Marshal(courses)
	finalJson, err := json.MarshalIndent(courses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonData := []byte(`
		{
			"coursename": "GO",
                "Price": 299,
                "website": "yt",
                "tags": ["web","backend"]	
        }
	`)

	var courses course
	checkValid := json.Valid(jsonData)

	if checkValid {
		fmt.Println("JSON is valid")
		json.Unmarshal(jsonData, &courses)
		fmt.Printf("%#v\n", courses)
	} else {
		fmt.Println("JSON is not valid..")
	}

	// when you want add data to key value

	var myData map[string]interface{}
	json.Unmarshal(jsonData, &myData)
	fmt.Printf("%#v\n", myData)

	for k, v := range myData {
		fmt.Printf("Key is %v and value is %v and Type of %T \n", k, v, v)
	}
}
