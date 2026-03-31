package main

import (
	"fmt"
	"slices"
)

func main() {
	// slice
	marks := []int{1, 2, 3, 5, 6, 7, 7, 9}
	fmt.Println(marks[len(marks)-1])
	fmt.Println(len(marks))
	fmt.Println(cap(marks))

	marks = append(marks, 34)
	fmt.Println(marks)

	
	slices.Reverse(marks) // reverses inplace brother, marks will be changed
	fmt.Println(marks)
	
	fmt.Println("----------------------------------")
	for i, val := range marks {
		fmt.Println(i, val)
	}

}
