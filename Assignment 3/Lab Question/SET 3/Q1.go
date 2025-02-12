//1. WAP in go language to accept two matrices and display it's multiplication.

package main

import "fmt"

func main() {
	var r1, c1, r2, c2 int

	fmt.Print("Enter rows and columns of first matrix: ")
	fmt.Scan(&r1, &c1)
	fmt.Print("Enter rows and columns of second matrix: ")
	fmt.Scan(&r2, &c2)

	if c1 != r2 {
		fmt.Println("Matrix multiplication not possible!")
		return
	}

	matrix1 := make([][]int, r1)
	matrix2 := make([][]int, r2)
	result := make([][]int, r1)

	for i := range matrix1 {
		matrix1[i] = make([]int, c1)
	}
	for i := range matrix2 {
		matrix2[i] = make([]int, c2)
	}
	for i := range result {
		result[i] = make([]int, c2)
	}

	fmt.Println("Enter elements of first matrix:")
	for i := 0; i < r1; i++ {
		for j := 0; j < c1; j++ {
			fmt.Scan(&matrix1[i][j])
		}
	}

	fmt.Println("Enter elements of second matrix:")
	for i := 0; i < r2; i++ {
		for j := 0; j < c2; j++ {
			fmt.Scan(&matrix2[i][j])
		}
	}

	for i := 0; i < r1; i++ {
		for j := 0; j < c2; j++ {
			for k := 0; k < c1; k++ {
				result[i][j] += matrix1[i][k] * matrix2[k][j]
			}
		}
	}

	fmt.Println("Resultant Matrix:")
	for _, row := range result {
		fmt.Println(row)
	}
}
