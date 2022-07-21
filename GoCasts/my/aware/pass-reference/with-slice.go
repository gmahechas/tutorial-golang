package main

import "fmt"

func main() {
	mySlice := []string{"a", "b", "c"}
	updateSlice(mySlice)
	fmt.Println(mySlice)
}

func updateSlice(s []string) {
	s[0] = "x"
}
