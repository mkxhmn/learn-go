package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Parry",
		contact: contactInfo{
			email:   "jim@email.com",
			zipCode: 9400,
		},
	}

	fmt.Printf("%v\n",jim)
}
