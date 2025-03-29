package main

// PaymentProcessor interface - Basic structure that will provide LSP
type PaymentProcessor interface {
	Pay(amount float64) string
}
