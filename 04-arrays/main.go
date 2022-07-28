package main

import "fmt"

func main() {
	nums := [5]int{1, 2, 3, 4, 5} // array, inmutable
	/* nums[5] = 6 // error */
	fmt.Println(nums)

	names := [4]string{"John", "Paul", "George", "Ringo"}
	fmt.Printf("%+v\n", names)

	coins := [...]int{1, 5, 10, 25, 50, 60, 100} // dinamyc length
	fmt.Printf("%T\n", coins)

	stringCoins := [...]string{0: "Penny", 1: "Nickel", 2: "Dime"}
	fmt.Printf("%+v\n", stringCoins)
}
