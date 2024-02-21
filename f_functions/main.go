package main

import "fmt"

func main() {
	var grettingsFirstPart string = getGrettingsFirst("Hello")
	var grettingsSecondPart string = getGrettingsSecond("World")

	fmt.Println(grettingsFirstPart + " " + grettingsSecondPart)
}

func getGrettingsFirst(grettings string) string {
	return grettings
}

func getGrettingsSecond(grettings string) string {
	return grettings
}
