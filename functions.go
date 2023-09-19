package main

import "fmt"

/*
// basic structure
func function_name (parameter type) return type {
	fucntion body
}
// parameter type and return type are optional
*/

func email_address(first_name, last_name string) {
	fmt.Println(first_name + last_name + "@abc.edu")
}
func addition(num1, num2 int) int { // here we need to mention return type of fucntion because we are explicitly using return statement
	return num1 + num2
}
func subtraction(val1, val2 int) { // here there is no return type because I am not returning anything
	result := val1 - val2
	fmt.Println(result)
}

func main() {
	email_address("laskhman", "rao")
	subtraction((addition(12, 12)), 20)
	message := func() { // assigning a funcion to a variable
		fmt.Println("Now we are good with functions")
	}
	message()
}
