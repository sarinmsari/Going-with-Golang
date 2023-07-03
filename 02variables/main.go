package main

import "fmt"

func main() {
	var username string = "Doops"
	fmt.Println(username)
	fmt.Printf("Variable is of type %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type %T \n", isLoggedIn)

	var smallValue uint8 = 252
	fmt.Println(smallValue)
	fmt.Printf("Variable is of type %T \n", smallValue)
	
	var smallFloat float32 = 252.159487263159
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type %T \n", smallFloat)
	
	//default values
	var defaultInt int
	fmt.Println(defaultInt)
	fmt.Printf("Variable is of type %T \n", defaultInt)
	
	//implicit type
	var noTypeVariable = "Hello Sarin"
	fmt.Println(noTypeVariable)
	fmt.Printf("Variable is of type %T \n", noTypeVariable)

	//no var style
	name:="Sarin"
	fmt.Println(name)
}
