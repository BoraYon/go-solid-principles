package main

import "fmt"

func processPayment(p PaymentMethod, amount float64) {
	fmt.Println(p.Pay(amount))
}

func main() {
	cc := CreditCard{CardNumber: "1234567890123456"}
	bt := BankTransfer{IBAN: "TR00001234123412341234"}

	processPayment(cc, 150.00)
	processPayment(bt, 950.00)

	fmt.Println("---------------------------------------------------")

	cp := CryptoPayment{
		WalletAddress: "BitcoinWalletAddress",
		Currency:      "BTC",
	}

	processPayment(cp, 0.015)

	// 🔒 OCP Applied:
	// We extended the system with a new payment type (CryptoPayment)
	// without modifying the existing payment types (CreditCard, BankTransfer)
	// or the core logic (processPayment). This adheres to the Open/Closed Principle.
}
