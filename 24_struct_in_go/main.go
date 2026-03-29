package main

import "fmt"

// defining a struct
type User struct {
	Name   string
	Age    int
	Gender string
	Salary float32
}

// https://www.notion.so/tactical-advantage/Go-33130115a66280dc94b7ccb214420b7c?source=copy_link#33230115a662801ca3a9db57c5d5052c
func main() {

	// creating a struct

	me := User{
		Name:   "Srikant V S",
		Age:    29,
		Gender: "M",
	}

	fmt.Printf("me: %v\n", me)
	// accessing
	fmt.Println("Age ", me.Age)


	// lets change some data, Struct fields are mutable
	me.Salary=100.55
	fmt.Println("Salary ", me.Salary)

}
