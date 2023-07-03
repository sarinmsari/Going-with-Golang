package main

import "fmt"

func main() {
	fmt.Println("welcome to loops in golang")

	days := []string{"sunday", "monday", "tuesday", "wednesday", "thursday", "friday", "saturday"}
	fmt.Println(days)

	/* for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	} */

	/* for i := range days {
		fmt.Println(days[i])
	} */

	for index, value := range days {
		fmt.Printf("%v is day %v\n", value, index+1)
	}

}
