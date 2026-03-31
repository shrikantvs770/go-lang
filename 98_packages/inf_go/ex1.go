package main

import (
	"fmt"

	"gopkg.in/inf.v0"
)

func main() {

	// https://www.notion.so/tactical-advantage/Go-33130115a66280dc94b7ccb214420b7c?source=copy_link#33630115a662801081d7fb25059b6f91
	num1 := inf.NewDec(123456798798, 1)
	fmt.Printf("num1: %v\n", num1)



	// okay so all the arithmetic stuff thik hai na..
	a1 := inf.NewDec(10000, 0)
	a2 := inf.NewDec(25000,0)

	result := new(inf.Dec)
	result = a1.Mul(a1,a2)

	fmt.Printf("result: %v\n", result)
}