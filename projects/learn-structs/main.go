package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	benjamin := person{firstName: "Benjamin", lastName: "Anderson"}
	fmt.Println(benjamin)
}
