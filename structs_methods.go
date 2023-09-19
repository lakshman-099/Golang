package main

import "fmt"

// Define the Account struct with corrected field names.
type Account struct {
	name    string
	balance float32
}

// Define methods for the Account struct.
func (a *Account) depositMoney(amount int) {
	a.balance += float32(amount)
}

func (a *Account) withdrawMoney(amount int) {
	a.balance -= float32(amount)
}

func (a Account) checkBalance() float32 {
	return a.balance
}

func main() {
	// Create two account instances.
	a1 := Account{"Lakshman", 5000}
	a2 := Account{"Ashok", 1000}

	fmt.Printf("Checking balance for %s: %.2f\n", a1.name, a1.checkBalance())
	fmt.Printf("%s balance before deposit: %.2f\n", a2.name, a2.checkBalance())
	a2.depositMoney(300)
	fmt.Printf("%s balance after deposit: %.2f\n", a2.name, a2.checkBalance())
}
