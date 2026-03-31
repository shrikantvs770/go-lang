package main

import "fmt"

func main() {
	type User struct {
		name    string
		address string
	}


	fmt.Printf("User: %v\n", User{})

	name := "srk"
	address:="ind"

	// if both key and value are same then no need to define multiple times
	fmt.Printf("%v\n", User{name, address})

	u2:= User{
		name: name,
		address: address,
	}

	fmt.Println(u2)


}
