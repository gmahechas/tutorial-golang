package main

import "fmt"

func main() {
	/* HelloWorld */
	/* 	fmt.Print("Hello")
	   	fmt.Print("World\n") */

	/*
		Hello
		World
	*/
	/* 	fmt.Println("Hello")
	   	fmt.Println("World") */

	/* s: string, d: int, f: float, v: value*/
	/* 	name := "Analuna"
	   	age := 3
	   	fmt.Printf("Hello %s you are %d\n", name, age) */

	/* 	msg := fmt.Sprintf("Hello %s you almost are %d\n", "Ana", 4)
	   	fmt.Println(msg) */

	name := "Analuna"
	fmt.Printf("name is: %T\n", name)

	fmt.Print("enter your name: ")
	fmt.Scanln(&name)

	fmt.Println("Hello", name)
}
