package main

import "fmt"

func main(){
 
	// 
	studentMap := make(map[string]int, 10) // 10 is the size here
	studentMap["srikant"] = 99
	studentMap["elias"] = 77
	studentMap["blair"]=67

	
	fmt.Printf("studentMap: %v\n", studentMap)
	fmt.Printf("studentMap size : %v\n", len(studentMap))



}