package main

import (
	"fmt"
)

func main() {

	var even []int = []int{2,4,8}
	odd := []int{1,2,3}

	var nos []int
	// := cannot use for assignment only for declaring variables
	// = only for assigning to existing variables
	
	nos=append(even, odd...)
	nos1:=append(even, odd...)

	fmt.Println(odd, even)
	fmt.Println(nos)
	fmt.Println(nos1)

}
