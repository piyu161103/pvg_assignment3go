//1. Find the Largest and Smallest Number in an Array

package main

import "fmt"

func main() {
	arr := []int{34, 78, 12, 45, 89, 3, 56}
	min, max := arr[0], arr[0]

	for _, num := range arr {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	fmt.Println("Smallest Number:", min)
	fmt.Println("Largest Number:", max)
}
