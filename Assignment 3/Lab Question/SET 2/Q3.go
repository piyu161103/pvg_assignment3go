//3. Student Details with Total and Average

package main

import "fmt"

type Student struct {
	RollNo   int
	Name     string
	Marks    [3]float64
	Total    float64
	Average  float64
}

func main() {
	var n int
	fmt.Print("Enter the number of students: ")
	fmt.Scan(&n)

	students := make([]Student, n)

	for i := 0; i < n; i++ {
		fmt.Printf("Enter details for student %d:\n", i+1)
		fmt.Print("Roll No: ")
		fmt.Scan(&students[i].RollNo)
		fmt.Print("Name: ")
		fmt.Scan(&students[i].Name)

		var sum float64
		for j := 0; j < 3; j++ {
			fmt.Printf("Enter mark %d: ", j+1)
			fmt.Scan(&students[i].Marks[j])
			sum += students[i].Marks[j]
		}
		students[i].Total = sum
		students[i].Average = sum / 3
	}

	fmt.Println("\nStudent Details:")
	for _, student := range students {
		fmt.Printf("Roll No: %d, Name: %s, Total: %.2f, Average: %.2f\n", student.RollNo, student.Name, student.Total, student.Average)
	}
}
