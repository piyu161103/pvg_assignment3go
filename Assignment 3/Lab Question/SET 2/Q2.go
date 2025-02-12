//2. Sort Array Elements in Ascending Order

package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{34, 12, 78, 56, 89, 3, 45}
	sort.Ints(arr)
	fmt.Println("Sorted Array in Ascending Order:", arr)
}
