package main

import "fmt"

// Declaring constants
const (
	pi       = 3.14159
	language = "Go"
)

func main() {
	// Using constants
	fmt.Printf("The value of pi is: %f\n", pi)
	fmt.Printf("Welcome to %s programming language!\n", language)

	// Enumerated constants
	const (
		Monday = iota
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		Sunday
	)

	// Using enumerated constants
	fmt.Println("Days of the week:")
	fmt.Println("Monday:", Monday)
	fmt.Println("Tuesday:", Tuesday)
	fmt.Println("Wednesday:", Wednesday)
	fmt.Println("Thursday:", Thursday)
	fmt.Println("Friday:", Friday)
	fmt.Println("Saturday:", Saturday)
	fmt.Println("Sunday:", Sunday)
}
