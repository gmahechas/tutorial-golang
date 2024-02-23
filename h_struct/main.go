package main

import (
	"fmt"
	"h_struct/user"
)

func main() {
	user1 := user.New(
		"tavogus",
		"tavogus@me.com",
		32,
		true,
	)
	user1.Print()
	newEmail := user1.GetEmail()
	fmt.Printf("new email: %s\n", newEmail)
	newEmail = "analuna@me.com"
	fmt.Printf("new email: %s\n", newEmail)

	user1.SetEmail(newEmail)
	user1.Print()

	user2 := user.New(
		user1.Username,
		user1.GetEmail(),
		user1.Age,
		user1.Active,
	)
	user2.Print()
}
