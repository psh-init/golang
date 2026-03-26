package main

import "fmt"

type BankAccount struct {
	name    string
	balance int
}

// deposit method
func (b *BankAccount) deposit(amount int) {
	b.balance = b.balance + amount
	fmt.Printf("%d deposited\n", amount)
}

// withdraw method
func (b *BankAccount) withdraw(amount int) {
	if amount > b.balance {
		fmt.Println("Insufficient balance")
		return
	}
	b.balance = b.balance - amount
	fmt.Printf("%d withdraw\n", amount)
}
func (b BankAccount) showBalance() {
	fmt.Println("Current balance:", b.balance)
}

func main() {

	acc := BankAccount{
		name:    "Priyanshu",
		balance: 1000,
	}

	acc.showBalance()
	acc.deposit(500)
	acc.withdraw(200)
	acc.showBalance()
}
