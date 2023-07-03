package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("slices on golang")

	var fruitlist = []string{"apple", "orange", "tomoto"}
	fmt.Printf("type of fruits is %T\n", fruitlist)

	fruitlist = append(fruitlist, "mango", "guava")
	fmt.Println(fruitlist)

	// fruitlist = append(fruitlist[1:3])
	// fmt.Println(fruitlist)

	fmt.Println(sort.StringsAreSorted(fruitlist))
	sort.Strings(fruitlist)
	fmt.Println(fruitlist)
	fmt.Println(sort.StringsAreSorted(fruitlist))

	var cources = []string{"js", "reactjs", "ts", "html/css", "go"}
	fmt.Println(cources)
	index := 3
	cources = append(cources[:index], cources[index+1:]...)
	fmt.Println(cources)
}
