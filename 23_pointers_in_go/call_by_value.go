package main

import "fmt"

func main() {

	a := 10
	fmt.Printf("a: %v\n", a) // 10
	m1(a)
	fmt.Printf("a: %v\n", a) // 10

	m2(&a)
	fmt.Printf("a: %v\n", a) // 30

}

func m1(a int) {
	a = 30
}

func m2(a *int){
	*a = 30
}
