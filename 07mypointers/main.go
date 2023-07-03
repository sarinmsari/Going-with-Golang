package main

import "fmt"

func main() {
	fmt.Println("class of pointers")

	var ptr *int
	fmt.Println("value of pointer is ", ptr)

	number := 50
	var ptr2 = &number
	fmt.Println("address pointed by the pointer is ", ptr2)
	fmt.Println("value of the pointer is ", *ptr2)

	*ptr2 = *ptr2 * 2
	fmt.Println("new vlaue of number is ", number)
}
