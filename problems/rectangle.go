package main

import "fmt"

// Rectangle represents a rectangle with length and width.
type Rectangle struct {
	length float64
	width  float64
}

// area calculates the area of a rectangle.
func (r Rectangle) area() float64 {
	return r.length * r.width
}

// perimeter calculates the perimeter of a rectangle.
func (r Rectangle) perimeter() float64 {
	return 2 * (r.length + r.width)
}

func main() {
	// Create an initial rectangle with the specified dimensions
	initialRectangle := Rectangle{length: 5.0, width: 3.0}

	// Calculate and print the area and perimeter of the initial rectangle
	fmt.Println("Initial rectangle")
	fmt.Println("Area:", initialRectangle.area())
	fmt.Println("Perimeter:", initialRectangle.perimeter())

	// Double the size of the rectangle
	initialRectangle.length *= 2
	initialRectangle.width *= 2

	// Calculate and print the area and perimeter of the doubled rectangle
	fmt.Println("\nDoubled rectangle")
	fmt.Println("Area:", initialRectangle.area())
	fmt.Println("Perimeter:", initialRectangle.perimeter())
}
