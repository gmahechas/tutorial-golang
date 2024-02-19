package main

import "fmt"

func main() {
	x := 10
	{
		x := 20 // bad practice, because it shadows the outer variable
		fmt.Println(x)
	}
	fmt.Println(x)
}
