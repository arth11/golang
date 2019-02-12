package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	contactInfo
}
type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	jim :=
		person{
			firstName: "jim",
			lastName:  "Party",
			contactInfo: contactInfo{
				email:   "ddd@gmaial.com",
				zipCode: 43434,
			},
		}
	jim.print()
}

func (p person) print() {
	fmt.Printf("%v", p)
}
