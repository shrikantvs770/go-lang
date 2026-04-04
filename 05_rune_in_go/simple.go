package main

import "fmt"

func main() {
	var a rune = '😀' // a rune
	fmt.Printf("a: %v %U\n", a, a)
	fmt.Println(len(string(a))) // takes 4 Bytes brother.
	fmt.Println(int32(a))// 128512


	var b int32 = 128512 // is same as var b rune = 128512
	fmt.Println(string(b)) // 😀


	
} 