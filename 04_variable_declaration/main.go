package main

import (
	"fmt"

)

func main(){
	var username string = "Srikant V S"

	// shortcut way, uses inference, but limited to a function
	title := "Sheeranahalli"
	age, gender := 29, 'M'


	fmt.Println(username, title, age, string(gender))

}