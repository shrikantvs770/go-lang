package main

import "fmt"

type Person struct {
	name   string
	gender string
}

type Teacher struct {
	Person // internally its like Person Person
	id     int
	salary float32
}

type Student struct {
	*Person
	school string
}

func main() {
	// https://www.notion.so/tactical-advantage/Go-33130115a66280dc94b7ccb214420b7c?source=copy_link#33730115a6628035b958f03779bac0c3
	t1 := Teacher{
		Person: Person{name: "John Doe", gender: "M"},
		id:     22,
		salary: 7800.88,
	}

	fmt.Printf("t1: %v\n", t1)
	// We can access the person fields directly
	fmt.Printf("salary %v\n", t1.salary)

	s1 := Student{
		Person: &Person{name: "Jane Doe", gender: "F"},
		school: "APS",
	}

	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s1: %v\n", s1.name) // auto dreferencing in action
	fmt.Printf("s1: %v\n", (*s1.Person).name)

	s1.name = "Elias"
	fmt.Printf("s1: %v\n", s1.name) // auto dreferencing in action

	(*s1.Person).name = "John"
	fmt.Printf("s1: %v\n", s1.name) // auto dreferencing in action

}
