package main

import (
	"fmt"
	"time"
)

type HTTPLogger struct{}

func (h *HTTPLogger) Send(entry LogEntry) {
	// Simulate sending an HTTP request
	time.Sleep(50 * time.Millisecond)
	fmt.Printf("SENT TO SERVER: [%s] %s\n", entry.Timestamp.Format(time.RFC3339), entry.Message)
}
