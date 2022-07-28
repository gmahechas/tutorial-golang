package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	Person
	salary int
}

func main() {
	p1 := Person{"Bob", 20}
	fmt.Println(p1)

	em1 := Employee{Person: Person{name: "Analuna", age: 30}, salary: 100}
	fmt.Printf("%+v\n", em1)
}
