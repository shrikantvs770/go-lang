package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
    Name  string `json:"name"`
    Age   int    `json:"age"`
}

func main() {
	body := `{"name": "Srikant", "age": 25}`

	var u User
	
	fmt.Printf("%p\n", &u) // That's an address
	json.Unmarshal([]byte(body), u)

	fmt.Println(u.Name) 
}