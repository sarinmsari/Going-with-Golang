package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Tutor    string
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("JSON in go 1")
	EncodeJson()
}

func EncodeJson() {
	lcoCources := []course{
		{"ReactJs", 499, "Hitesh Choudary", "abc123", []string{"web", "js", "frontend"}},
		{"NodeJs", 699, "Hitesh Choudary", "bcd123", []string{"web", "js", "backend"}},
		{"GO lang", 899, "Hitesh Choudary", "cde123", nil},
	}

	finalJson, err := json.MarshalIndent(lcoCources, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(finalJson))
	DecodeJson(finalJson)
}

func DecodeJson(JsonData []byte) {
	checkValid := json.Valid(JsonData)
	fmt.Println(string(JsonData))
	if checkValid {
		fmt.Println("valid json")

		/* var lcoCourse []course
		json.Unmarshal(JsonData, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse) */

		var myJsonData []map[string]interface{}
		json.Unmarshal(JsonData, &myJsonData)
		fmt.Printf("%#v\n", myJsonData)

	} else {
		fmt.Printf("INVALID JSON")
	}
}
