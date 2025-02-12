package main

import (
	"fmt"
)

func main() {
	var emptySlice []int
	fmt.Println("Empty slice:", emptySlice)
	var sliceArray []int
	fmt.Println("Declared slice array:", sliceArray)
	sliceLiteral := []int{10, 20, 30, 40, 50}
	fmt.Println("Slice initialized with literal:", sliceLiteral
	sliceLiteral = append(sliceLiteral, 60, 70)
	fmt.Println("After appending elements:", sliceLiteral)
	fmt.Println("Length of the slice:", len(sliceLiteral))
	fmt.Println("Capacity of the slice:", cap(sliceLiteral))
	subSlice := sliceLiteral[2:5]
	fmt.Println("Sub-slice from index 2 to 4:", subSlice)
	fmt.Println("Iterating through slice elements:")
	for i, value := range sliceLiteral {
		fmt.Printf("Index %d: %d\n", i, value)
	}
}
