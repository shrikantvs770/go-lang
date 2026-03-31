package main

import (
	"fmt"
	"math/big"
)

func main() {



	a1:=big.NewInt(1000)
	a2:=big.NewInt(100)

	res:= new(big.Int)

	res = res.Add(a1, a2).Sub(a1, a2) // 1100, will be overwritten and finally 1000-100 = 900

	fmt.Printf("res: %v\n", res)

	// true chaining

	res1 := new(big.Int)
	res1 = res1.Add(a1, a2).Sub(res1, a2)
	fmt.Printf("res1: %v\n", res1)



}