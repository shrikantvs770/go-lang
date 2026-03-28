package main

import "fmt"

func main() {

	day := 38

	switch day {
	case 1, 2, 3:
		fmt.Println("Hola")
	case 4, 5:
		fmt.Println("Hello")
	default:
		fmt.Println("Unknown Gunman")
	}

}
