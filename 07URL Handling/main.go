package main

import (
	"fmt"
	"net/url"
)

const myURL string = "https://lco.dev:3000/learn?coursename=reactjs&payment=sdfnxncaosifei"

func main() {
	fmt.Println("Welcome to URL handling in go")
	fmt.Println(myURL)

	//parsing
	result, _ := url.Parse(myURL)
	fmt.Println("url scheme: ", result.Scheme)
	fmt.Println("url host: ", result.Host)
	fmt.Println("url path: ", result.Path)
	fmt.Println("url port: ", result.Port())
	fmt.Println("url rewquery: ", result.RawQuery)

	qparams := result.Query()
	fmt.Printf("Type of qparams is %T\n", qparams)
	fmt.Println("coursename: ", qparams["coursename"])
	for _, val := range qparams {
		fmt.Println("param is ", val)
	}

	partsofurl := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/learn",
		RawQuery: "user=sarin",
	}
	fullurl := partsofurl.String()

	fmt.Println(fullurl)
}
