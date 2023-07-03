package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("files in golang")
	content := "This is a content for a file"

	file, err := os.Create("./myGoFile.txt")
	checkNilError(err)

	length, err := io.WriteString(file, content)
	checkNilError(err)
	fmt.Println("length is ", length)
	defer file.Close()
	readFile("./myGoFile.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilError(err)
	fmt.Println("data inside the file is ", string(databyte))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
