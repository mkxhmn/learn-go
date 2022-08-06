package main

import "fmt"

func main() {
	colors := make(map[string]string)

	//var colors map[string]string

	//colors := map[string]string{
	//	"red":   "#ff0000",
	//	"green": "#333",
	//}

	// add values into map
	colors["white"] = "#fff"
	colors["red"] = "#ff0000"

	// delete function
	delete(colors, "white")

	fmt.Println(colors)
}
