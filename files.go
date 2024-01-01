// ioutil.WriteFile is deprecated: As of Go 1.16

package main

import (
	"fmt"
	"os"
)

// writeToFile writes the given content to a file with the specified filename.
func writeToFile(filename string, content string) {
	// Create or open the file
	file, err := os.Create(filename)
	checkError(err)
	defer file.Close()                 // Deferring the closing of the file to ensure it's closed at the end of this function
	_, err = file.WriteString(content) // Write content to the file
	checkError(err)
	fmt.Println("File written successfully")
	// Even though we deferred file.Close() above, it will be executed just before this function returns
}

// readFromFile reads the content of the file with the specified filename and prints it.
func readFromFile(filename string) {
	// Open the file
	file, err := os.Open(filename)
	checkError(err)
	defer file.Close()
	data, err := os.ReadFile(filename)
	checkError(err)
	fmt.Println(string(data)) // Print the content read from the file
}

// checkError is a helper function to handle errors.
func checkError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
}

func main() {
	filename := "myfile.txt"
	writeToFile(filename, "This is new content in this repo")
	readFromFile(filename)
}
