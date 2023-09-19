package main

import "fmt"

func sum(arr [4]int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

func main() {
	// Declare an array of strings with 3 elements
	names := [3]string{"Lakshman", "Anudeep", "Purna"}

	// Declare an empty array without specifying its length explicitly
	emptyList := [...]int{}

	// Declare a 2D array of integers
	matrix := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	// Declare an array of integers with an initializer list
	intList := [...]int{23, 45, 56, 44}

	// Print the empty list and the sum of intList
	fmt.Println(emptyList)
	fmt.Println(sum(intList))

	// Update the second element of the names array
	names[1] = "XYZ"

	// Print the modified names array and the matrix
	fmt.Println(names[1])
	fmt.Println(matrix)
}
