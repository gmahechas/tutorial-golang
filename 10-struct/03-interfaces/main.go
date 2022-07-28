package main

import "fmt"

type Dog struct{}
type Fish struct{}
type Bird struct{}

type Animal interface {
	Speak()
}

func Speak(animal Animal) {
	animal.Speak()
}

func (dog *Dog) Speak() {
	fmt.Println("Woof!")
}

func (fish *Fish) Speak() {
	fmt.Println("Blub!")
}

func (bird *Bird) Speak() {
	fmt.Println("Tweet!")
}

func main() {
	dog := Dog{}
	fish := Fish{}
	bird := Bird{}

	Speak(&dog)
	Speak(&fish)
	Speak(&bird)
}
