package main

import (
	"fmt"
	"math"
)

func main() {
	// Variable declaration and initialization
	var age int
	fmt.Println("Age:", age)

	// Assigning values to a variable
	age = 22
	fmt.Println("Age:", age)

	// Declaring a variable with an initial value
	var ageWithInitialValue int = 22
	fmt.Println("Age:", ageWithInitialValue)

	// Short hand declaration
	count := 10
	fmt.Println("Count =", count)

	// Short hand declaration for multiple variables
	nameShort, ageShort := "Mandeep", 22
	fmt.Println("Name:", nameShort, "age ", ageShort)

	// Runtime value assignment
	aRuntime, bRuntime := 182.8, 563.8
	cRuntime := math.Min(aRuntime, bRuntime)
	fmt.Println("Minimum value is", cRuntime)
}
