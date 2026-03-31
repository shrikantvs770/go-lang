// Demonstrates nested structs and returning a pointer to a struct.
// Shows how to access nested fields and use pointers with structs.
package main

import "fmt"

type Wrestler struct {
	name   string
	weight int
}

type WWE struct {
	wrestler Wrestler
}

func m1(w Wrestler) *WWE {

	w.name = "Oba Femi"
	w.weight = 300

	company := WWE{w}
	return &company
}

// nested struct Demo

func main() {

	w1 := m1(Wrestler{})

	// or see how we add add * saying its pointer, but the above syntax we don't even see if its pointer or not ha ha.

	// var w2 *WWE = m1(Wrestler{})

	var w2 = m1(Wrestler{})  // this is also valid we don't know if this is pointer or not but it is.

	fmt.Println(w1.wrestler)
	fmt.Println(w1.wrestler.name)
	fmt.Println(w1.wrestler.weight)

	fmt.Println(w2)

}
