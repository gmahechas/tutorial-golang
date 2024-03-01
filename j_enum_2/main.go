package main

import "fmt"

func main() {
	enumOne()
	enumTwo()
}

func enumOne() {
	var name *string
	// var name2 *string = new(string) // this initializes name2 with a pointer to a string ""

	newName := "tavogus"
	name = &newName

	if name != nil {
		fmt.Println(*name) // Imprime "tavogus"
	} else {
		fmt.Println("name is nil")
	}
}

func enumTwo() {
	p := Person{name: "tavogus", age: nil}
	p.GetAge()
}

type Person struct {
	name string
	age  *int
}

func (p Person) GetAge() {
	if p.age == nil {
		fmt.Println("age is nil")
	} else {
		fmt.Println(p.age)
	}
}
