package main

import "fmt"

// Function without parameters and return value
func sayHello() {
	fmt.Println("Hello, World!")
}

// Function with parameters and return value
func addNumbers(a, b int) int {
	sum := a + b
	return sum
}

// Function with multiple return values
func calculateRectangleArea(length, width float64) (float64, float64) {
	area := length * width
	perimeter := 2 * (length + width)
	return area, perimeter
}

func main() {
	// Calling a function without parameters and return value
	sayHello()

	// Calling a function with parameters and return value
	result := addNumbers(10, 20)
	fmt.Println("Sum:", result)

	// Calling a function with multiple return values
	area, perimeter := calculateRectangleArea(5.0, 3.0)
	fmt.Printf("Area: %f, Perimeter: %f\n", area, perimeter)
}
