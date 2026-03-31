package main

import "fmt"

type User struct {
	Name string
	Age int
}

func main(){
	// i:=new(int)

	var temp = new(User)
	temp.Name = "Elias"
	temp.Age = 33

	// another way

	temp2:= &User{
		Name: "randy",
		Age: 78,
	}

	// or

	var temp3 *User = &User{ // bhai type to automatically pata chalta hai na?
		Name: "randy",
		Age: 78,
	}


// see below yahi ho raha hai
	x:=1
	p:=&x

	fmt.Printf("temp: %v %T\n", temp, temp)
	fmt.Printf("temp: %v %T\n", *temp, *temp)

	fmt.Printf("temp2: %v\n", temp2)

}