package main

import "fmt"

// -------------------------------
// Interfaces
// -------------------------------

// Requires Display()
type Displayer interface {
	Display()
}

// Requires Update(string)
type Updater interface {
	Update(string)
}

// -------------------------------
// Struct
// -------------------------------
type Employee struct {
	Name string
	Age  int
}

// VALUE RECEIVER
// - Works on a COPY of Employee
// - Safe, does NOT modify original
func (emp Employee) Display() {
	fmt.Printf("emp: %v\n", emp)
}

// POINTER RECEIVER
// - Works on ORIGINAL object
// - Can modify state
func (emp *Employee) Update(newName string) {
	emp.Name = newName
}

func main() {

	// =========================================
	// CASE 1: VALUE VARIABLE
	// =========================================

	e1 := Employee{Name: "Elias", Age: 22}

	// Direct call → value receiver
	e1.Display()

	// Pointer method on VALUE
	// Go rewrites:
	// e1.Update("John Cena") → (&e1).Update("John Cena")
	// Works because e1 is addressable
	e1.Update("John Cena")

	// Update modified original (because pointer receiver)
	fmt.Printf("e1: %v\n", e1)

	// =========================================
	// CASE 2: POINTER VARIABLE
	// =========================================

	var e2 *Employee
	e2 = &Employee{Name: "Bianca", Age: 56}

	// Pointer calling value method
	// Go rewrites:
	// e2.Display() → (*e2).Display()
	e2.Display()

	// Pointer calling pointer method → direct
	e2.Update("Trish")

	fmt.Printf("e2: %v\n", *e2)

	// =========================================
	// NOW THE TRUTH: INTERFACES (STRICT)
	// =========================================

	var d Displayer
	var u Updater

	// ---- Displayer ----
	d = e1   // ✅ works
	d = e2   // ✅ works

	// WHY?
	// Display() is VALUE receiver
	// → belongs to Employee e1
	// → also usable by *Employee e2


	// ---- Updater ----

	// u = e1 // ❌ DOES NOT COMPILE
	// ERROR:
	// Employee does NOT implement Updater
	// (Update method has pointer receiver)

	u = e2 // ✅ works

	// WHY?
	// Update() belongs ONLY to *Employee and e2 is pointer

	u.Update("Interface Call")
	fmt.Printf("after interface update: %v\n", *e2)


	// =========================================
	// FAILURE CASES (IMPORTANT)
	// =========================================

	// ---- Case 1: Non-addressable value ----

	getEmployee().Update("X") // ❌ WILL NOT COMPILE

	// WHY?
	// getEmployee() returns a VALUE (temporary)
	// Go CANNOT do:
	// (&getEmployee()).Update(...) ❌ illegal


	// ---- Case 2: Map value ----

	m := map[string]Employee{
		"a": {Name: "A", Age: 10},
	}

	// m["a"].Update("X") // ❌ WILL NOT COMPILE

	// WHY?
	// map values are NOT addressable
	// so Go cannot take &m["a"]


	// ---- Case 3: Interface strictness ----

	var emp Employee

	// var up Updater = emp // ❌

	// WHY?
	// Interface does NOT allow:
	// auto &emp
	// Must match EXACT method set


	// =========================================
	// HELPER FUNCTION
	// =========================================
}

func getEmployee() Employee {
	return Employee{Name: "Temp"}
}