package main

import "fmt"

type BankTransfer struct {
	IBAN string
}

func (b BankTransfer) Pay(amount float64) string {
	return fmt.Sprintf("Transferred %.2f via bank to IBAN %s", amount, b.IBAN)
}
