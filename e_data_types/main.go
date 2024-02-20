package main

import "fmt"

func main() {
	// integers
	const a int = 10
	const b int = -10
	fmt.Println(a, b)

	// floats
	const c float64 = 10.0
	const d float64 = -10.0
	fmt.Println(c, d)

	// booleans
	const e bool = true
	fmt.Println(e)

	// char
	const f rune = 'a'
	fmt.Println(f)

	// tuples
	const g, h = 10, 20
	fmt.Println(g, h)
	const i, j = 10.0, true
	fmt.Println(i, j)
	const k, l = true, "hello"
	fmt.Println(k, l)

	// arrays or vectors
	var m [3]int = [3]int{1, 2, 3}
	fmt.Println(m)

	// slices, no number inside []
	var n []int = []int{1, 2, 3}
	fmt.Println(n)

	o := make([]int, 3)
	fmt.Println(o)

	// strings
	const p string = "hello"
	fmt.Println(p)
}
