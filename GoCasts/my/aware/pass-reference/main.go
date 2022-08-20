package main

import "fmt"

func main() {
	name := "Bill"
	updateValue(&name)
	fmt.Println(name) // Bill
}

func updateValue(name *string) {
	fmt.Println(name) // 0xc000..
	// name = "Alex" // error
	*name = "Alex"
}
