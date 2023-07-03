package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	fmt.Println("post form request")
	performPostForm()
}

func performPostForm() {
	const myUrl = "http://localhost:8000/postform"

	//formdata
	data := url.Values{}
	data.Add("firstname", "hitesh")
	data.Add("lastname", "choudary")
	data.Add("email", "hitesh@lco.com")

	response, err := http.PostForm(myUrl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}
