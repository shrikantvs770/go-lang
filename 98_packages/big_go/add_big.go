package main

import (
	"fmt"
	"math/big"
)

func main() {
	// 100-digit numbers as strings
	num1Str := "6680208701343756102724374235382590136529299288974514710267653027885426397749684404416985645988385615"
	num2Str := "1872794801195728574515739386552637485630186385894688046247070975766246364286097323836486592176447576"

	// Initialize big.Int variables
	n1 := new(big.Int)
	n2 := new(big.Int)

	// Set values from strings (base 10)
	n1.SetString(num1Str, 10)
	n2.SetString(num2Str, 10)

	// Add the numbers
	sum := new(big.Int).Add(n1, n2)

	fmt.Println("Sum:", sum.String())
}
