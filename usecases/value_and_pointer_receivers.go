package main

import "fmt"

// Employee is a simple struct
type Employee struct {
	Name string
	Age  int
}

// VALUE RECEIVER
// This method gets a COPY of Employee
// Any modification here DOES NOT affect original
func (emp Employee) Display() {
	fmt.Printf("emp: %v\n", emp)
}

// POINTER RECEIVER
// This method gets a POINTER to Employee
// Any modification here WILL affect original
func (emp *Employee) Update(newName string) {
	emp.Name = newName
}

func main() {

	// -------------------------------
	// CASE 1: VALUE VARIABLE
	// -------------------------------
	e1 := Employee{Name: "Elias", Age: 22}

	// Calling value receiver → direct call
	e1.Display()

	// Calling pointer receiver on VALUE
	// Go automatically converts:
	// e1.Update(...) → (&e1).Update(...)
	e1.Update("John Cena")

	// Since Update uses pointer receiver,
	// original struct is modified
	fmt.Printf("e1: %v\n", e1)

	// -------------------------------
	// CASE 2: POINTER VARIABLE
	// -------------------------------
	var e2 *Employee
	e2 = &Employee{Name: "Bianca", Age: 56}

	// Calling value receiver on POINTER
	// Go automatically converts:
	// e2.Display() → (*e2).Display()
	e2.Display()

	// Calling pointer receiver → direct call
	e2.Update("Trish")

	fmt.Printf("e2: %v\n", *e2)
}