/*
the basic structure for infinite loop

	for {
		//......
		if you want to break use break
	}

we no need to mention any condition after for
*/

package main

import "fmt"

func main() {

	counter := 0

	// Create an infinite loop
	for {

		fmt.Println("This is an infinite loop")

		// Increment the counter
		counter++

		// If the counter is greater than or equal to 10, break out of the loop
		if counter >= 10 {
			break
		}
	}
}
