package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("POST request with GO frontend")
	performPostJsonRequest()
}

func performPostJsonRequest() {
	const myurl = "http://localhost:8000/post"

	//fake json data
	requestBody := strings.NewReader(`
		{
			"coursename":"let's go with golang"
			"price": 0
			"tutor": "Hitesh Choudary"
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
