package main

import "fmt"

func main()  {

	result := sumAll(1,2,3,4,5,6,7,8,9)
	fmt.Printf("result: %v\n", result)


	// that's a slice brother
	result = sumAll([]int{1,2,3,4,5,6,7,8,9}...)
	fmt.Printf("result2: %v\n", result)

}


func sumAll(nums  ...int) int {
	
	sum:=0
	for _,val :=range(nums) {
		sum+=val
	}

	return  sum
}