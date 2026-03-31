package main

import "fmt"

type Animal interface {
	Speak() string // Method and return type
	Move() string
	AnimalName() // only method
}

type Dog struct {
	Name string
}

func (d *Dog) Speak() string {
	return "Woof"
}

func (d *Dog) Move() string {
	return "runs"
}

func (d *Dog) AnimalName() {
	fmt.Printf("d: %v\n", *d)
}
func main() {

	// var animal Animal = Dog{Name: "Tom"} // ERROR::: as method belongs to *Dog not Dog
	// https://www.notion.so/tactical-advantage/Go-33130115a66280dc94b7ccb214420b7c?source=copy_link#33730115a6628047be79dafb4c77552d
	var animal Animal = &Dog{Name: "Tom"}

	fmt.Println(animal.Move())
	fmt.Println(animal.Speak())
	animal.AnimalName()

	fmt.Printf("type of animal : %T\n", animal)
	fmt.Printf("animal: %v\n", *animal) // ERROR invalid operation: cannot indirect animal
	// it doesn't work because animal is not a pointer but an interface and inside that there is a pointer, so lets take it out


	
	d1:=animal.(*Dog)
	fmt.Printf("d1: %v\n", *d1)

}
