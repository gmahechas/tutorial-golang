package main

import (
	"fmt"
	"strings"
)

func repeat(n int) func(chain string) string {
	return func(chain string) string {
		return strings.Repeat(chain, n)
	}
}
func main() {
	repeat := repeat(3)
	fmt.Println(repeat("*"))
	/* 	func() {
		fmt.Println("Hello")
	}() */
}
