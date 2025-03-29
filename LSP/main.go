package main

import "fmt"

func processPayment(processor PaymentProcessor, amount float64) {
	fmt.Println(processor.Pay(amount))
}

func main() {
	cc := CreditCardPayment{CardNumber: "1234567890123456"}
	cp := CryptoPayment{WalletAddress: "1BitcoinWalletExample123", Currency: "BTC"}
	fd := FraudDetectionService{BlockedCards: map[string]bool{"1234567890123456": true}}

	fmt.Println("Processing payments...")
	processPayment(cc, 150.00)
	processPayment(cp, 0.015)

	// Fraud check
	fmt.Println("Performing fraud check...")
	if fd.IsFraud(cc.CardNumber) {
		fmt.Println("Transaction blocked: Fraud detected on card ending with", cc.CardNumber[len(cc.CardNumber)-4:])
	} else {
		processPayment(cc, 150.00)
	}
}
