package main

import (
	"fmt"
	"strings"
)

func main() {
	const myHello string = "World"
	const grettings string = "Hello" + " " + myHello
	fmt.Println(grettings)

	var myString strings.Builder
	const first = "Hello "
	const second = "World"
	fmt.Fprint(&myString, first, second, "!")
	myMessage := myString.String()
	fmt.Println(myMessage)
}
