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

func (p person) greet() {
	fmt.Printf("hello! my name is %v %v. Nice to meet you\n", p.firstName, p.lastName)
}

// a pointer that points at the person
func (p *person) updateName(newName string) {
	// whenever we place *, that pointer is turn into a value
	(*p).firstName = newName
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Parry",
		contactInfo: contactInfo{
			email:   "jim@email.com",
			zipCode: 9400,
		},
	}

	jim.greet()

	// turn to memory address or pointer
	jimPointer := &jim
	jimPointer.updateName("Jimmy")
	jim.greet()
}
