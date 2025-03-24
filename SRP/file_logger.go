package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

type FileLogger struct {
	file *os.File
	mu   sync.Mutex
}

func NewFileLogger(path string) *FileLogger {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return &FileLogger{file: f}
}

func (f *FileLogger) Write(entry LogEntry) {
	f.mu.Lock()
	defer f.mu.Unlock()
	fmt.Fprintf(f.file, "[%s] %s\n", entry.Timestamp.Format(time.RFC3339), entry.Message)
}
