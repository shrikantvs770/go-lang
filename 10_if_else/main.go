package main

import "fmt"

func main() {
	age := 45
	if age > 18 {
		fmt.Println("Person can vote")
	} else {
		fmt.Println("Person cannot vote")
	}

	// if else ladder
	marks := 60
	if marks > 90 && marks < 100 {
		fmt.Println("A")
	} else if marks > 80 && marks < 90 {
		fmt.Println("B")
	} else {
		fmt.Println("Fail")
	}

}
