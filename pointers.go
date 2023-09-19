/*
   In Go, "&" is used to create a pointer to a variable,
   and "*" is used to dereference a pointer.
*/

package main

import (
	"fmt"
)

func main() {
	// Create an integer variable and set its initial value to 7
	originalValue := 7

	// Create a pointer that points to the memory location of originalValue
	pointerToValue := &originalValue

	// Dereference the pointer to get the value it points to
	dereferencedValue := *pointerToValue

	// Modify the value through the dereferenced copy (does not affect originalValue)
	dereferencedValue = 8

	// Print the original value (remains unchanged)
	fmt.Println("Original Value:", originalValue)

	// Print the modified dereferenced value
	fmt.Println("Dereferenced Value:", dereferencedValue)
}
