package main

import "fmt"

// CreditCardPayment struct
type CreditCardPayment struct {
	CardNumber string
}

func (c CreditCardPayment) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f using Credit Card ending with %s", amount, c.CardNumber[len(c.CardNumber)-4:])
}
