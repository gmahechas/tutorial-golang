package main

import "fmt"

func main() {
	// borrowing
	var a int = 10
	var b int = a
	fmt.Println(&a)
	fmt.Println(&b)
}
