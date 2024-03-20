package main

import "fmt"

func main() {
	var n []int = []int{1, 2}
	n = append(n, 3)
	fmt.Println(n)

	//hashset
	var o map[int]bool = map[int]bool{1: true, 2: true}
	o[3] = true
	fmt.Println(o)

	// hashmap/map, there is no duplicates
	var m map[string]int = map[string]int{"a": 1, "b": 2}
	m["c"] = 3
	fmt.Println(m)
}
