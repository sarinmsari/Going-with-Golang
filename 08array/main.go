package main

import "fmt"

func main() {
	fmt.Println("array in golang")

	var fruitList [4]string
	fruitList[0] = "apple"
	fruitList[1] = "orange"
	fruitList[3] = "peach"

	fmt.Println("fruitlist is ", fruitList)

}
