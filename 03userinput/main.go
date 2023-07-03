package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name: ")

	//comma ok || comma err
	input, _ := reader.ReadString('\n')

	welcome := "Hello " + input + "Welcome to the go programming"
	fmt.Println(welcome)
}
