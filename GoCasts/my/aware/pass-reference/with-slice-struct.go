package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	mySlice := []person{{"John", "Doe"}, {"Jane", "Doe"}, {"Jim", "Party"}}
	/* updateSlice(mySlice) */
	updateSingle(mySlice[0])
	fmt.Println(mySlice)
}

func updateSlice(s []person) {
	s[0].firstName = "x"
}

func updateSingle(p person) {
	p.firstName = "x"
}
