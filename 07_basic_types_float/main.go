package main

import "fmt"

func main(){
	var isAdult bool
	isLogged := true
	fmt.Printf("isAdult: %v\n", isAdult)
	fmt.Printf("isLogged: %v\n", isLogged)

	fmt.Println("AND", isAdult && isLogged)
	fmt.Println("OR", isAdult || isLogged)


}
