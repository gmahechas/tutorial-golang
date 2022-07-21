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
	/* contactInfo contactInfo */
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *person) updateName(newFirstName string) {
	fmt.Printf("&a: %p\n", &p)
	/* p.firstName = newFirstName */
	(*p).firstName = newFirstName
}
