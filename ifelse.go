package main

import "fmt"

func main() {
	// Example of if-else statement
	number := 15

	if number%2 == 0 {
		fmt.Printf("%d is even.\n", number)
	} else {
		fmt.Printf("%d is odd.\n", number)
	}

	// Example of if-else if-else statement
	grade := 85

	if grade >= 90 {
		fmt.Println("A")
	} else if grade >= 80 {
		fmt.Println("B")
	} else if grade >= 70 {
		fmt.Println("C")
	} else {
		fmt.Println("D")
	}

	// Example of if statement with a short variable declaration
	if x := 42; x > 0 {
		fmt.Println("x is positive")
	} else if x == 0 {
		fmt.Println("x is zero")
	} else {
		fmt.Println("x is negative")
	}
}
