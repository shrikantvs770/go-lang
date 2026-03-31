package main

import "fmt"

func main() {
	i:=new(int) // i is not pointer to int
	fmt.Printf("i: %v\n", i)
	fmt.Printf("i: %v\n", *i)
}