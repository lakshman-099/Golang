/*
basic structure for while loop
for i <10 {
	//......
}
there is no seperate keyword for while loop
you just have to use for as a keyword
in while loop any condition you mention after for
the variable in condition must be initialized before the loop
*/

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var guess int
	// Generate a random secret number between 0 and 99
	secretNumber := rand.Intn(100)
	found := false

	// Use a loop to repeatedly prompt the user for guesses until they guess correctly
	for !found {
		fmt.Println("Enter a number (between 0 and 99): ")
		fmt.Scan(&guess)

		if guess == secretNumber {
			fmt.Printf("Congratulations! You guessed the number correctly: %d\n", guess)
			found = true
		} else if guess < secretNumber {
			fmt.Println("Your guess is too low. Try again.")
		} else if guess > secretNumber {
			fmt.Println("Your guess is too high. Try again.")
		}
	}
}
