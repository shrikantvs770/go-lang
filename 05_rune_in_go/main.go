package main

import (
	"fmt"
	"reflect"
)

func main() {

	
	uniq := "ħëļļò" // what' the length? is it 5?
	fmt.Println("uniq len : ", len(uniq)) // says 10
	fmt.Printf("uniq unicode: %U\n", []rune(uniq)[0]) // U+0127


	// to get correct length convert into rune slice
	fmt.Println("uniq corr len : ", len([]rune(uniq)))         //5


	fmt.Println(len("é")) // says 2

	fmt.Println([]rune("é")) // [233]
	fmt.Println(rune('å')) // 229
	fmt.Println(len([]rune("å"))) // 1



	for i, r := range "hello"  { // reflect.TypeOf(r) its int32 means its a rune brother, also 
		fmt.Println(i,reflect.TypeOf(r))
	}

	// 	Behind the scenes:
	// Reads bytes
	// Decodes UTF-8
	// Returns:
	// i → byte index
	// r → decoded rune
	
}