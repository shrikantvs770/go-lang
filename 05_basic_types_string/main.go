package main

import "fmt"
import "strings"

func main(){
	firstName := "SRIKANT"
	lastName := "V S"

	fmt.Println(firstName+" "+lastName) // concatenation
	fmt.Println(strings.ToLower(firstName))

	kaPlace := strings.Index(firstName, "KA")
	fmt.Println("KA occurs at", kaPlace)
}
