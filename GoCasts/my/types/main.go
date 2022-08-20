package main

import (
	"fmt"
	"reflect"
)

type namesType []string

func (names namesType) changeFirst() {
	names[0] = "Paul McCartney"
}

func main() {
	names := namesType{"John", "Paul", "George"}
	names.changeFirst()
	fmt.Println(names)
	fmt.Println(reflect.TypeOf(names))
}
