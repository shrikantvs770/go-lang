package main

import (
	"fmt"
	"strings"
)

type User struct {
	Name string
	Age  int
}

func printUser(user *User) {
	fmt.Printf("user: %v\n", *user)
	user.Captalize()
	fmt.Printf("user: %v\n", *user)
	fmt.Printf("user: %v\n", *user.Captalize())
	// fmt.Printf("user: %v\n", *&user.Captalize()) // error
	fmt.Printf("user: %v\n", user)
	fmt.Printf("user: %v\n", &user)
	// fmt.Printf("user: %v\n", &user.Captalize()) // error
}

func (user *User) Captalize() {
	user.Name = strings.ToUpper(user.Name)
}

func main() {

	user := User{
		Name: "John",
		Age:  25,
	}
	printUser(&user)
}
