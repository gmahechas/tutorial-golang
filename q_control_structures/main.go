package main

import "fmt"

func main() {
	// if
	var number int = 3

	if number < 5 {
		fmt.Println("condition was true")
	} else if number == 5 {
		fmt.Println("condition was false")
	} else {
		fmt.Println("condition was false")
	}

	// ternary
	result := 3
	if result > 5 {
		result = 7
	} else {
		result = 2
	}
	fmt.Println("result:", result)

	// while loop
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	// for
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}
