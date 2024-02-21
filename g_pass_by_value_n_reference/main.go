package main

import "fmt"

func main() {
	passByValue()
	fmt.Println("#######")
	passByReference()
}

func passByValue() {
	var x int = 10
	fmt.Printf("The initial value of x is: %d\n", x)
	passByValueOne(x)
	fmt.Printf("The final value of x is: %d\n", x)
}

func passByValueOne(x int) {
	x = x + 1
	fmt.Printf("The value of x in the inner scope is: %d\n", x)
}

func passByReference() {
	var x int = 10
	fmt.Printf("The initial value of x is: %d\n", x)
	passByReferenceOne(&x)
	fmt.Printf("The final value of x is: %d\n", x)
}

func passByReferenceOne(x *int) {
	*x = *x + 1
	fmt.Printf("The value of x in the inner scope is: %d\n", *x)
}
