package main

// single import
import "time"

// multiple imports
import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("5^2 is : ", math.Pow(5, 2))
	fmt.Println("sqrt(25) : ", math.Sqrt(25))
	fmt.Println("Today date is : ", time.Now().UTC())

}
