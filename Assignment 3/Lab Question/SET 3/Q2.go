//2. Employee Details with Maximum Salary

package main

import "fmt"

type Employee struct {
	ID     int
	Name   string
	Salary float64
}

func main() {
	var n int
	fmt.Print("Enter the number of employees: ")
	fmt.Scan(&n)

	employees := make([]Employee, n)

	for i := 0; i < n; i++ {
		fmt.Printf("Enter details for employee %d:\n", i+1)
		fmt.Print("ID: ")
		fmt.Scan(&employees[i].ID)
		fmt.Print("Name: ")
		fmt.Scan(&employees[i].Name)
		fmt.Print("Salary: ")
		fmt.Scan(&employees[i].Salary)
	}

	maxSalary := employees[0]
	for _, emp := range employees {
		if emp.Salary > maxSalary.Salary {
			maxSalary = emp
		}
	}

	fmt.Println("Employee with Maximum Salary:")
	fmt.Printf("ID: %d, Name: %s, Salary: %.2f\n", maxSalary.ID, maxSalary.Name, maxSalary.Salary)
}
