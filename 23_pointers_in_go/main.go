package main

import "fmt"

// https://www.notion.so/tactical-advantage/Go-33130115a66280dc94b7ccb214420b7c?source=copy_link#33230115a66280fe9afcd340daef8b08

func main() {

	x := 10
	p := &x
	var q *int = &x // another way to declare a pointer brother

	fmt.Println("x", x) // 10
	fmt.Println("p", p) // 0xc0000120d8
	fmt.Println("*p", *p)
	fmt.Println("q", q)
	fmt.Println("*q", *q)

	*p = 20
	fmt.Println(x) // 20, because p is pointer to x

}
