// these are addition to arrays
// these are dynamically sized arrays and slices are preferred way to work with arrays in Go

package main

import "fmt"

// max returns the maximum value in a slice of integers.
func max(slice []int) int {
	maxValue := 0
	for i := 0; i < len(slice); i++ {
		if slice[i] > maxValue {
			maxValue = slice[i]
		}
	}
	return maxValue
}

func main() {
	// Declare a slice of integers.
	sliceA := []int{2, 3, 4, 5}

	// Append an element to the slice.
	sliceA = append(sliceA, 10)
	fmt.Println(sliceA)

	// Create a new slice 'sliceB' with a length of 5 and copy elements from 'sliceA'.
	sliceB := make([]int, 5)
	copy(sliceB, sliceA)
	fmt.Print(sliceB)
	fmt.Printf("\n%T", sliceB)

	// Append 'sliceB' elements to 'sliceA' and store the result in 'sliceC'.
	sliceC := append(sliceA, sliceB...)
	fmt.Print("\n", sliceC)

	// Find and print the maximum value in 'sliceC'.
	maxValue := max(sliceC)
	fmt.Print("\n", maxValue)
}
