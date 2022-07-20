package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	benjamin := person{firstName: "Benjamin", lastName: "Anderson"}
	fmt.Println(benjamin)

	var alex person
  alex.firstName="Alex"
  alex.lastName="Billy"

  // print format
  fmt.Printf("%+v", alex)
}
