package main

// Correct FraudDetection struct
type FraudDetectionService struct {
	BlockedCards map[string]bool
}

func (fd *FraudDetectionService) IsFraud(cardNumber string) bool {
	return fd.BlockedCards[cardNumber]
}
