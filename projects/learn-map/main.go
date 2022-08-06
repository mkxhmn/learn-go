package main

import "fmt"

func main() {
	//colors := make(map[string]string)

	//var colors map[string]string

	colors := map[string]string{
		"green": "#333",
	}

	// add values into map
	colors["white"] = "#fff"
	colors["red"] = "#ff0000"

	// delete function
	delete(colors, "white")

	printMap(colors)

	fmt.Println(colors)
}

func printMap(c map[string]string) {
	// iterate map
	for color, hex := range c {
		fmt.Printf("%v with hex code %v\n", color, hex)
	}
}
