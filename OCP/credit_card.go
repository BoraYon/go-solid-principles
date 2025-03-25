package main

import "fmt"

type CreditCard struct {
	CardNumber string
}

func (c CreditCard) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f using Credit Card ending with %s", amount, c.CardNumber[len(c.CardNumber)-4:])
}
