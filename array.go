package main

import "fmt"

func main() {
	// Declaring and initializing an array
	var numbers [5]int
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	numbers[3] = 4
	numbers[4] = 5

	// Accessing elements of the array
	fmt.Println("Array:", numbers)
	fmt.Println("First element:", numbers[0])
	fmt.Println("Second element:", numbers[1])
	fmt.Println("Third element:", numbers[2])

	// Initializing an array with values
	names := [3]string{"Alice", "Bob", "Charlie"}

	// Iterating through the array
	fmt.Println("\nNames:")
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// Using the range keyword to iterate through the array
	fmt.Println("\nUsing Range:")
	for index, value := range names {
		fmt.Printf("Index: %d, Value: %s\n", index, value)
	}

	// Multidimensional array
	var matrix [2][3]int
	matrix[0][0] = 1
	matrix[0][1] = 2
	matrix[0][2] = 3
	matrix[1][0] = 4
	matrix[1][1] = 5
	matrix[1][2] = 6

	// Accessing elements of a multidimensional array
	fmt.Println("\nMatrix:")
	fmt.Println(matrix[0])
	fmt.Println(matrix[1])
}
