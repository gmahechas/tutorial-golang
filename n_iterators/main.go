package main

import "fmt"

func main() {
	// array
	array := [3]int{1, 2, 3}

	for index, value := range array {
		fmt.Println(index, value)
	}

	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}

	// slice
	slice := []int{4, 5, 6}

	for index, value := range slice {
		fmt.Println(index, value)
	}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	// map
	myMap := map[string]int{
		"a": 1,
		"b": 2,
	}

	for key, value := range myMap {
		fmt.Println(key, value)
	}
}
