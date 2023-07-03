package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("web request")

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close() //callers responsibility to colse the connection

	fmt.Printf("response is of type %T", response)

	databytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(databytes))
}
