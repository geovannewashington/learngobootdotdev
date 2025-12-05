package main

import (
	"fmt"
)

type UserAccount struct {
	firstName string
	lastName  string
	balance   float64
}

// reciver, the special argument that comes before the method name
func (u *UserAccount) setBalance(newBalance float64) {
	u.balance = newBalance 
}

func (u *UserAccount) getBalance() {
	fmt.Printf("balance: %f\n", u.balance)
	fmt.Printf("===============================\n")
}

func main() {
	// creates an instance of a user account on the bank
	user01 := UserAccount{
		firstName: "John",
		lastName: "Doe",
		balance: 4000.00,
	}
	user01.getBalance()	
	
	// set new balance
	user01.setBalance(5000.00)

	// see change
	user01.getBalance()	
}
