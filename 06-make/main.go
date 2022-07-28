package main

import "fmt"

func main() {
	nums := make([]int, 0, 3)
	fmt.Printf("len: %v, Cap %v, %p\n", len(nums), cap(nums), nums)
	fmt.Println(nums)
	nums = append(nums, 1, 2, 3)
	fmt.Println(nums)
	nums = append(nums, 4, 5, 5)
	fmt.Printf("len: %v, Cap %v, %p\n", len(nums), cap(nums), nums)
	fmt.Println(nums)
}
