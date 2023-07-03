package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study")

	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01/02/2006 Monday 15:04"))

	createdDate := time.Date(2023, time.June, 28, 15, 04, 0, 0, time.UTC)
	fmt.Println(createdDate.Format("Created on 02/01/2006 15:04"))
}
