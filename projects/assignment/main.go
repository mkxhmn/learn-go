package main

import "fmt"

func main() {
	// assignment one
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("assignment one: odd or even")
	oddOrEven(numbers)

	// assignment two
	s := square{sideLength: 10}
	t := triangle{height: 10, base: 10}

	fmt.Println("\nassignment two: area")
	printArea(s)
	printArea(t)

}
