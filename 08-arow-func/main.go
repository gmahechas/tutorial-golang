package main

import "fmt"

func main() {
	/* 	func() {
		fmt.Println("Hello")
	}() */
	myFunc := func() {
		fmt.Println("Hello")

	}
	myFunc()
}
