package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (person *Person) speak() {
	fmt.Println("Hello, my name is", person.name)
}

func (person *Person) changeName(newName string) {
	person.name = newName
}

func main() {
	/* 	p1 := Person{"Bob", 20}
	   	fmt.Println(p1)
	   	p1.name = "Analuna" // it works with/without pointer
	   	fmt.Println(p1)
	   	p1.speak() */

	p2 := Person{name: "Alice", age: 30}
	fmt.Printf("%+v\n", p2)
	p2.changeName("Isabella")
	fmt.Printf("%+v\n", p2)
	p2.speak()
}
