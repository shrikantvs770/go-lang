package main

import (
	"fmt"
	"math/big"
)
func main() {
	a1 := big.NewInt(103298000006500000)
	a2 := big.NewInt(33)

	res := new(big.Int)
	res = a1.Add(a1, a2)
	fmt.Printf("res: %v\n", res)

}