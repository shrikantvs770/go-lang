package main

import "fmt"

func main() {
	var myarr [][]int = [][]int{
		{1, 2, 2, 3, 6},
		{1, 2, 3, 77},
		{1, 2, 3, 5}, // note the extra ,
	}

	myarr2:= [...][]int{{1,2},{3,4}}

	fmt.Println(myarr)
	fmt.Println(myarr2)
}
