package main

import "fmt"

func main() {
	var grade string
	fmt.Print("Enter your grade(A,B,C,D or F): ")
	fmt.Scan(&grade)
	// Switch case statement to handle different grades.
	// Fallthrough is used to continue execution to the next case without evaluating its condition.
	switch grade {
	case "A":
		fmt.Println("Excellent!")
		fallthrough
	case "B":
		fmt.Println("Good job!")
		// No fallthrough here, so when "B" is matched, execution doesn't continue to the next case.
	case "C":
		fmt.Println("You Passed")
	case "D", "F":
		fmt.Println("You need to improve")
	default:
		fmt.Println("Invalid Grade")
	}
}
