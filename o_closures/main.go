package main

import "fmt"

func main() {
	lambda()
}

func lambda() {
	sum := func(x int, y int) int {
		return x + y
	}
	fmt.Println(sum(5, 2))
}
