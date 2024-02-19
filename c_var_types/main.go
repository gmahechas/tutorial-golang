package main

import "fmt"

func main() {
	const MaxPoints float64 = 100.0 // inmutable, it's exported bacause it starts with a capital letter
	const minSize float64 = 50.0    // inmutable, it's not exported
	Truth := true                   // implicit type
	fmt.Println(MaxPoints)
	fmt.Println(minSize)
	fmt.Println(Truth)

	const (
		MyConstOne = 5
		MyConstTwo = MyConstOne + 7
	)
	fmt.Println(MyConstOne, MyConstTwo)

	const (
		North = iota // 0
		East         // 1
		South        // 2
		West         // 3
	)
	fmt.Println(South)

	var a int = 10
	a = a + 1
	fmt.Printf("%d\n", a)
}
