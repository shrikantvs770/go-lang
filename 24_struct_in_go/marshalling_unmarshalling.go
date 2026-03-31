package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"nameXX"` // caps is important in Name
	Age int `json:"AgeXX"`
}


func main() {

	var u User = User{
		Name : "srikantvs26",
		Age : 29,
	}


	res, _ := json.Marshal(u)

	
	fmt.Printf("u: %v\n", u)
	fmt.Printf("res: %v\n", string(res))

	um := User{} // or var um User
	var data = "{\"nameXX\":\"srikantvs26\",\"AgeXX\":29}"
	json.Unmarshal([] byte(data), &um)
	fmt.Printf("um: %v\n", um)

}