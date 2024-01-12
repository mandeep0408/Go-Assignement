package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Example of basic switch usage
	fingerExample()

	// Example of a switch with a default case
	defaultCaseExample()

	// Example of multiple expressions in a case
	multipleExpressionsExample()

	// Example of an expressionless switch
	expressionlessSwitchExample()

	// Example of switch with fallthrough
	fallthroughExample()

	// Example of breaking out of a switch
	breakExample()

	// Example of generating a random even number and breaking the outer loop
	randomEvenNumberExample()
}

func fingerExample() {
	selectedFinger := 3                                  // Change the value
	fmt.Printf("Selected Finger %d is ", selectedFinger) // Change the variable name
	switch selectedFinger {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	case 5:
		fmt.Println("Pinky")
	}
}

func defaultCaseExample() {
	finger := 8
	fmt.Printf("Finger %d is ", finger) // Change the variable name
	switch finger {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	case 5:
		fmt.Println("Pinky")
	default:
		fmt.Println("Incorrect finger number")
	}
}

func multipleExpressionsExample() {
	selectedLetter := "o"                                  // Change the value
	fmt.Printf("Selected Letter %s is a ", selectedLetter) // Change the variable name
	switch selectedLetter {
	case "a", "e", "i", "o", "u":
		fmt.Println("vowel")
	default:
		fmt.Println("not a vowel")
	}
}

func expressionlessSwitchExample() {
	numValue := 75 // Change the value
	switch {
	case numValue >= 0 && numValue <= 50:
		fmt.Printf("%d is greater than 0 and less than 50\n", numValue) // Change the variable name
	case numValue >= 51 && numValue <= 100:
		fmt.Printf("%d is greater than 51 and less than 100\n", numValue) // Change the variable name
	}
}

func fallthroughExample() {
	numValue := 75 // Change the value
	switch {
	case numValue < 50:
		fmt.Printf("%d is lesser than 50\n", numValue) // Change the variable name
		fallthrough
	case numValue < 100:
		fmt.Printf("%d is lesser than 100\n", numValue) // Change the variable name
	}
}

func breakExample() {
	numValue := -5 // Change the value
	switch {
	case numValue < 50:
		if numValue < 0 {
			break
		}
		fmt.Printf("%d is lesser than 50\n", numValue) // Change the variable name
		fallthrough
	case numValue < 100:
		fmt.Printf("%d is lesser than 100\n", numValue) // Change the variable name
	}
}

func randomEvenNumberExample() {
randloop:
	for {
		switch generatedNumber := rand.Intn(100); { // Change the variable name
		case generatedNumber%2 == 0:
			fmt.Printf("Generated even number %d\n", generatedNumber) // Change the variable name
			break randloop
		}
	}
}
