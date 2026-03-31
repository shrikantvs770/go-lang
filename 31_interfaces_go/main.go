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

func (d Dog) Speak() string {
	return "Woof"
}

func (d Dog) Move() string {
	return "runs"
}

func (d Dog) AnimalName() {
	fmt.Printf("d: %v\n", d)
}
func main() {

	var animal Animal = Dog{Name: "Tom"}
	fmt.Println(animal.Move())
	fmt.Println(animal.Speak())
	animal.AnimalName()
	fmt.Printf("animal: %v\n", animal)

}
