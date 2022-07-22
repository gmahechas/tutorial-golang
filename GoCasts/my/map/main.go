package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}

	/* var colors map[string]string */

	/* 	colors := make(map[int]string)
	   	colors[10] = "#ff0000" */

	printMap(colors)

	fmt.Println(colors)
}

func printMap(colors map[string]string) {
	for color, hex := range colors {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
