package main

import (
	"fmt"

)
type User struct {
	Name string
	Age int
}

func main(){
// https://www.notion.so/tactical-advantage/Go-33130115a66280dc94b7ccb214420b7c?source=copy_link#33530115a662803a8de2ea1fbfa90ba7
	var user User
	if user == (User{}){
		fmt.Println("its empty")
	}

	var user1 *User
	if user1 == nil {
		fmt.Println("This is empty too. lets allocate")
		user1 = &User{
			Name: "Elias",
			Age: 45,
		}
		fmt.Printf("user1: %v\n", *user1)
		fmt.Printf("user1: %v\n", user1)
		
	}

}