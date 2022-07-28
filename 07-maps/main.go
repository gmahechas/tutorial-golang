package main

import "fmt"

func main() {
	/* 	myMap := make(map[int]string)
	   	fmt.Println(myMap)

	   	myMap[1] = "one"
	   	fmt.Println(myMap)
	   	myMap[2] = "two"
	   	fmt.Println(myMap) */

	students := make(map[string][]int)
	fmt.Println(students)
	students["Bob"] = []int{1, 2, 3}
	students["Alice"] = []int{2, 3, 4}
	fmt.Println(students)
}
