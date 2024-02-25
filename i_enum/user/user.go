package user

import "fmt"

/* ####### user ####### */
type User struct {
	Username string
	Email    string
	Age      int
	Active   bool
	Role     Role
	Website  Website
}

func New(username, email string, age int, active bool, role Role, website Website) *User {
	return &User{
		Username: username,
		Email:    email,
		Age:      age,
		Active:   active,
		Role:     role,
		Website:  website,
	}
}

func (user *User) Print() {
	fmt.Printf("Username: %s, Email: %s, Age: %d, Active: %t, Role: %d, Website: %d\n", user.Username, user.Email, user.Age, user.Active, user.Role, user.Website)
}

/* ####### role ####### */
type Role int

const (
	ADMIN Role = iota
	USER
)

func CheckAccess(role Role) bool {
	return role == ADMIN
}

/* ####### website ####### */
type Website int

const (
	Github Website = iota
	Facebook
	Twitter
)
