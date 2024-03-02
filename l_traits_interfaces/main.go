package main

import "fmt"

type Speaks interface {
	Speak()
}

type Human struct{}
type Animal struct{}

func (h Human) Speak() {
	fmt.Println("Hello there!")
}

func (a Animal) Speak() {
	fmt.Println("I'm an animal")
}

// duck typing
func Speak(s Speaks) {
	s.Speak()
}

func main() {
	var human Human
	var animal Animal

	// this is implicitly implemented
	human.Speak()
	animal.Speak()

	// this is explicitly implemented
	Speak(human)
	Speak(animal)
}
