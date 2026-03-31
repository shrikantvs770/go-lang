package main

import (
	"fmt"
)

func main() {
	var foo = 'a'
	fmt.Printf("byte: %v\n", foo)


	var num int = 97
	fmt.Printf("num: %v\n", byte(num))
	fmt.Printf("num: %v\n", string(num))
	fmt.Printf("num: %v\n", string(byte(num)))
	fmt.Printf("num: %v\n", string(rune(num)))

	my := []byte("Hello") // real workhorse
	fmt.Println(my)



	// String to byte conversion
	b1 := []byte("hello")
	// Byte to String conversion 
	s1 := string(b1)
	fmt.Printf("s1: %v type is %T \n", s1, s1)



}
