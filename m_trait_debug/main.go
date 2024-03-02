package main

import "fmt"

type User struct {
	Username string
	Age      int
}

func (u User) String() string {
	return fmt.Sprintf("Username :): %s, Age: %d", u.Username, u.Age)
}

func main() {
	user := User{
		Username: "tavogus",
		Age:      32,
	}
	fmt.Print(user, "\n")
	fmt.Println(user)
	fmt.Printf("%+v\n", user) // %+v prints all fields
	fmt.Printf("%#v\n", user) // %#v prints the type
	fmt.Printf("%T\n", user)  // %T prints the type
	fmt.Printf("%v\n", user)  // %v prints the value (same as default Println)
}
