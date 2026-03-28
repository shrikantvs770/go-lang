package main

import (
	"fmt"
)

func main() {

	/*
		The capacity of the slice is equal to its length. A second integer argument may be provided to specify a different capacity; it must be no smaller than the length. For example, make([]int, 0, 10) allocates an underlying array of size 10 and returns a slice of length 0 and capacity 10 that is backed by this underlying array.
	*/

	marks1 := make([]int, 10)
	marks2 := make([]int, 0, 10)


	fmt.Println(len(marks1), cap(marks1), marks1)
	fmt.Println(len(marks2), cap(marks2), marks2)
	
	// marks2[4]=56 will not work because len is 0, lets use append
	marks2 = append(marks2, 45)
	marks1[3]=78

	fmt.Println(marks1, marks2)





	




}
