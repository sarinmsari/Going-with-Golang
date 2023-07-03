package main

import "fmt"

func main() {
	fmt.Println("struct in golang")

	Sarin := User{"Sarin", "hi@sarinm.tech", 21, true}
	fmt.Println(Sarin)
	fmt.Printf("Details of user: %+v\n", Sarin)
	fmt.Printf("Name of user is %v and email of user is %v\n", Sarin.Name, Sarin.Email)
}

type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}
