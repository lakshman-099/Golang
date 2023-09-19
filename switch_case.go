package main

import "fmt"

func main() {
	var age int
	fmt.Print("Enter the age: ")
	fmt.Scan(&age)
	// Call the age_classification function to determine the age category and store the result.
	classification := age_classification(age)
	// Print the age classification.
	fmt.Println(classification)
}

// age_classification takes an age as input and returns a string representing the age category.
func age_classification(age int) string {
	// Using a switch statement to classify age into different categories based on given conditions.
	switch {
	case age == 0:
		return "NewBorn"
	case age <= 3 && age >= 1:
		return "toddler"
	case age >= 4 && age <= 12:
		return "child"
	case age >= 13 && age <= 17:
		return "teenager"
	case age >= 18:
		return "adult"
	default:
		return "Invalid input"
	}
}
