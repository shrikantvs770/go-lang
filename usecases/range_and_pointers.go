package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func PrintUserDetails(user *User) {
	fmt.Printf("user: %v\n", user)
	fmt.Printf("user: %v %v\n", (*user).Name, (*user).Age)
	fmt.Printf("user: %v %v\n", (user).Name, (user).Age)
}

// https://www.notion.so/tactical-advantage/Go-33130115a66280dc94b7ccb214420b7c?source=copy_link#33830115a662804b85cdde23cb83f1ac
func main() {

	userList := []User{
		User{
			Name: "John",
			Age:  45,
		},

		User{
			Name: "Elias",
			Age:  67,
		},
	}

	// Go does not auto-dereference or auto-take address in range or function calls.
	// range gives a copy, so pass &slice[i] if the function needs a pointer.

	for _, v := range userList {
		PrintUserDetails(&v)
		fmt.Printf("-----------------------\n")
	}

	// fmt.Printf("userList: %v\n", userList)

}
