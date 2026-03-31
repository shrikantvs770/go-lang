package main

import (
	"fmt"
	"os"
)

func main(){
	homeEnv := os.Getenv("HOME")
	codeSpaceName := os.Getenv("CODESPACE_NAME")
	wwe := os.Getenv("WWE")// return"""

	wwe,ok := os.LookupEnv("WWE")

	os.Setenv("USERNAME","srikantvs")
	username := os.Getenv("USERNAME")

	fmt.Printf("homeEnv: %v\n", homeEnv)
	fmt.Printf("codeSpaceName: %v\n", codeSpaceName)
	fmt.Printf("username: %v\n", username)
	fmt.Printf("wwe: %v\n", wwe) // empty

	if !ok {
		fmt.Println("WWE is not defined")
	}

}