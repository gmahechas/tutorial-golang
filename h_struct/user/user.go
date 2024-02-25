package user

import "fmt"

type User struct {
	Username string
	email    string // private
	Age      int
	Active   bool
}

func New(username, email string, age int, active bool) *User {
	return &User{
		Username: username,
		email:    email,
		Age:      age,
		Active:   active,
	}
}

func (user *User) GetEmail() string {
	return user.email
}

func (user *User) Print() {
	fmt.Printf("Username: %s, Email: %s, Age: %d, Active: %t\n", user.Username, user.email, user.Age, user.Active)
}

func (user *User) SetEmail(newEmail string) {
	user.email = newEmail
}
