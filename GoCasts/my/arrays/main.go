package main

import (
	"fmt"
)

func main() {
	myArray := [2]string{"Hello", "World"}
	updateArray(&myArray)
	updateArrayItem(&myArray[1])
	// myArray[2] = "!!!" // can't add items
	fmt.Println(myArray)
	for i, s := range myArray {
		fmt.Println(i, s)
	}

}

func updateArray(array *[2]string) {
	array[0] = "Goodbye"
}

func updateArrayItem(strOne *string) {
	*strOne = "Planet"
}
