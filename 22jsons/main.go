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
	EncodeJson()
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
