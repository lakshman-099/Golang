/*
// basic syntax
for initialization; contion; increment {
	// ......
}
*/

package main

import "fmt"

func main() {
	// Let's create a for loop to find Prime numbers in a range
	start := 10
	end := 50

	//iterate through each number in specified range
	for num := start; num <= end; num++ {
		isPrime := true

		// check for factors of num from 2 to num/2
		// if num is divisible by any value in this range, it's not prime
		for i := 2; i <= num/2; i++ {
			if num%i == 0 {
				isPrime = false
				// exit the loop since num is not prime
				break
			}
		}

		// if isPrime is still true, num is a prime number
		if isPrime {
			fmt.Printf("%d ", num)
		}
	}
}
