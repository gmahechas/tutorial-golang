package main

import "fmt"

type Point[T int | float64] struct {
	x T
	y T
}

func main() {
	var pointInt Point[int] = Point[int]{x: 7, y: 20}
	fmt.Println(pointInt)

	var pointFloat Point[float64] = Point[float64]{x: 10.0, y: 7.0}
	fmt.Println(pointFloat)

	fmt.Println(calculateArea(pointInt))
	fmt.Println(calculateArea(pointFloat))
}

func calculateArea[T int | float64](point Point[T]) T {
	return point.x * point.y
}
