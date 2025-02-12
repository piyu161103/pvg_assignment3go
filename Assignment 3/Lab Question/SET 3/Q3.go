//3. Demonstrating Slice Operations

package main

import "fmt"

func main() {
	slice := []int{10, 20, 30, 40, 50}

	// Append
	slice = append(slice, 60)
	fmt.Println("After append:", slice)

	// Remove an element (index 2)
	index := 2
	slice = append(slice[:index], slice[index+1:]...)
	fmt.Println("After removal:", slice)

	// Copy to a new slice
	newSlice := make([]int, len(slice))
	copy(newSlice, slice)
	fmt.Println("Copied slice:", newSlice)
}
