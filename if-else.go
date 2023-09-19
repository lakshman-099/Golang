package main

import "fmt"

func main() {
	var score int
	fmt.Print("Enter the student's score: ")
	fmt.Scan(&score)

	// Determine and display the letter grade
	letterGrade := determineLetterGrade(score)
	fmt.Printf("Letter Grade: %s", letterGrade) // Include format specifier %s to indicate where to insert the value letterGrade
}

func determineLetterGrade(score int) string {
	// Write code to determine the letter grade using if-else conditions
	if score >= 90 {
		return "A"
	} else if score >= 80 && score <= 89 {
		return "B"
	} else if score >= 70 && score <= 79 {
		return "C"
	} else if score >= 60 && score <= 69 {
		return "D"
	} else {
		return "F"
	}
}
