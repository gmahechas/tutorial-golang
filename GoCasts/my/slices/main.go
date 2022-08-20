package main

import "fmt"

func main() {
	mySlice := []string{"Hello", "World"}
	updateSlice(mySlice)
	updateSliceItem(&mySlice[1])
	mySlice = append(mySlice, "!!!")
	fmt.Println(mySlice)
	for i, s := range mySlice {
		fmt.Println(i, s)
	}
}

func updateSlice(slice []string) {
	slice[0] = "Goodbye"
}

func updateSliceItem(strOne *string) {
	*strOne = "Planet"
}
