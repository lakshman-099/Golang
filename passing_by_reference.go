package main

import "fmt"

// modifyValue takes a pointer to an integer as a parameter.
func modifyValue(ptr *int) {
	*ptr = 10 // Modifying the value that ptr points to
}

func main() {
	x := 8
	fmt.Println("Before:", x)
	// This allows modifyValue to modify the original x variable.
	modifyValue(&x)

	fmt.Println("After:", x)
}
