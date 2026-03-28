package main

import (
	"fmt"
)

func main() {
	nos := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var evens []int
	for _, v := range nos { // _ we don't want to do anything with index so _ otherwise you can use i
		if v%2 == 0 {
			evens = append(evens, v)
		}
	}

	fmt.Println(evens)
}
