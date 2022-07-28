package main

import "fmt"

func main() {
	/* 	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	   	fmt.Println(nums)
	   	nums = append(nums, 11)
	   	fmt.Println(nums)

	   	fmt.Printf("len: %v, Cap %v, %p\n", len(nums), cap(nums), nums)
	   	fmt.Printf("%T\n", nums) */

	names := make([]string, 3, 5)
	fmt.Printf("%+v\n", names)
	fmt.Printf("len: %v, Cap %v, %p\n", len(names), cap(names), names)
}
