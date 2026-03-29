package main
import "fmt"

type Employee struct {
	EmpName string
	Department string
	Salary float32
}

func main(){
	employee := Employee {
		EmpName: "JohnCena",
		Department: "Locker Room",
		Salary: 9000.55,
	}

	fmt.Printf("employee: %v\n", employee)
	employee.Update()
	fmt.Println("After updating salary ", employee.Salary) // 9000.55, because its pass by value

	employee.Update2()
	fmt.Println("After pass by pointer ", employee.Salary) // After pass by pointer  10000.99
} 

// method value receiver, it will get copy and changes are local to the method, but not reflected
func (emp Employee) Update() {
	emp.Salary=10000.99
}

// method pointer receiver
func (emp *Employee) Update2() {
	emp.Salary = 10000.99
}