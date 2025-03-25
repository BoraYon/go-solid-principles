package main

import "fmt"

type CryptoPayment struct {
	WalletAddress string
	Currency      string
}

func (c CryptoPayment) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f %s to wallet %s", amount, c.Currency, c.WalletAddress)
}
