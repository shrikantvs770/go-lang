package main

import "fmt"

type Address struct {
	country string
}

type AddressMap map[string]*Address

func main() {

	p:= AddressMap{
		"johncena":&Address{country: "USA"},
		"bryan": &Address{country: "India"},
		"mvp": &Address{country: "France"},
	}
	// Now I have heard that range automatically does dereferenceing for us so no need to use * for Address
	for k,v := range(p) {
		fmt.Printf("k:v %v %v\n", k, v.country)
		
	}
}