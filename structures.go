/*
type StructName struct {
    Field1 DataType1
    Field2 DataType2
    // Add more fields as needed
}

*/

package main

import "fmt"

// Define a Book struct with lowercase field names (private)
type Book struct {
	title  string  // Title of the book
	author string  // Author of the book
	price  float32 // Price of the book
}

func main() {
	// Create instances of the Book struct
	book1 := Book{
		title:  "The Great Gatsby",
		author: "F. Scott Fitzgerald",
		price:  9.99,
	}

	// Create another Book instance without field names
	book2 := Book{
		"To Kill a Mockingbird",
		"Harper Lee",
		12.99,
	}

	// Create a slice to store multiple books
	books := []Book{book1, book2}

	// Print the list of books
	fmt.Println("List of Books:")
	for _, book := range books {
		fmt.Println("Title:", book.title)
		fmt.Println("Author:", book.author)
		fmt.Println("Price: $", book.price)
		fmt.Println()
	}
}
