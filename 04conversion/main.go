package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Thanks for your bussines with us")
	fmt.Println("Rate us from 1 to 5")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		i := 0
		fmt.Printf("Thanks for rating: ")
		for i = 0; i < int(numRating); i++ {
			fmt.Printf("*")
		}
		if numRating > float64(i) {
			fmt.Printf("+")
		}
	}
}
