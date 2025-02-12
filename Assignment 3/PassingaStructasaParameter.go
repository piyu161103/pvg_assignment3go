package main

import "fmt"
type Rectangle struct {
	Length int
	Width  int
}

func calculateArea(rect Rectangle) int {
	return rect.Length * rect.Width
}

func modifyDimensions(rect *Rectangle, newLength, newWidth int) {
	rect.Length = newLength
	rect.Width = newWidth
}

func main() {
	rect := Rectangle{Length: 10, Width: 5}

	area := calculateArea(rect)
	fmt.Println("Area of rectangle:", area)
	modifyDimensions(&rect, 20, 10)
	fmt.Println("Modified rectangle:", rect)
}
