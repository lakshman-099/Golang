// maps are unordered key-value pairs

package main

import "fmt"

func main() {
	// Declare and initialize a map with string keys and int values.
	serverCounts := make(map[string]int)
	serverCounts = map[string]int{"servers": 2, "load_balancers": 3, "gateways": 2, "data_centers": 5}

	// Declare and initialize another map with int keys and int values.
	myMap2 := map[int]int{1: 1, 2: 2}

	// Insert a new key-value pair into myMap2.
	myMap2[3] = 4

	// Delete a key from serverCounts.
	delete(serverCounts, "gateways")

	// Check the existence of a key in serverCounts.
	_, found := serverCounts["servers"]
	if found {
		fmt.Printf("The 'servers' key exists in the map.")
	}

	fmt.Print("\n", myMap2)

	// Iterate through serverCounts and print key-value pairs.
	for key, value := range serverCounts {
		fmt.Printf("\n%s : %d", key, value)
	}
}
