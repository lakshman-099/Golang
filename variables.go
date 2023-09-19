package main

import "fmt"

func main() {
	var myname = "lakshman"
	fmt.Println("my name is", myname)

	// type annotation
	var occupation string = "Associate"
	fmt.Println(myname, " occupation is ", occupation)

	// unintialized variable
	var sum int
	fmt.Println(myname, "salary is ", sum) // by default if we don't initialize value it's going to be zero

	// short variable declaration
	company := "WHO" // ':=' this notation automatically infer the variable data type based on assigned value
	fmt.Println(myname, "is working in ", company)
}
