package main

import "fmt"

func main()  {

	// 
	res := add(1,2,3)
	fmt.Println(res)


	// order matters here
	sum, sub, mul :=calcv1(10,20,30)
	fmt.Printf("mul: %v\n", mul)
	fmt.Printf("sub: %v\n", sub)
	fmt.Printf("add: %v\n", sum)

	// Okay I only need mul brother
	_, _, multpx := calcv1(10, 20, 30)
	fmt.Printf("multpx: %v\n", multpx)


}

// just a normal function brother, int is the return type, as a, b and c have same type so int at the end is enough
func add(a, b, c int) int {
	return a+b+c
}


// the add, sub, mul in the return type are kind of declaration, order matters to callee
func calcv1(a, b, c int) (sum, sub, mul int) {
	sum = a+b+c
	sub = a-b-c
	mul = a*b*c
	return // its called naked return brother
}