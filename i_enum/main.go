package main

import "i_enum/user"

func main() {
	user1 := user.New(
		"tavogus",
		"tavogus@me.com",
		32,
		true,
		user.ADMIN,
		user.Twitter,
	)
	user1.Print()
}
