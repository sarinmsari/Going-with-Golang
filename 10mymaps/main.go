package main

import "fmt"

func main() {
	fmt.Println("maps in golang")

	languages := make(map[string]string)

	languages["js"] = "JavaScript"
	languages["ts"] = "TypeScript"
	languages["py"] = "Python"

	fmt.Println("list of all languages: ", languages)
	fmt.Println("js: ", languages["js"])

	delete(languages, "py")
	fmt.Println(languages)

	for key, value := range languages {
		fmt.Println(key, value)
	}
}
