package main

import "fmt"

func main() {
	// Example of a for loop
	fmt.Println("For Loop:")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// Example of a while loop
	fmt.Println("\nWhile Loop:")
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	// Example of a range loop with an array
	fmt.Println("\nRange Loop:")
	numbers := [3]int{1, 2, 3}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
}
