package main

import (
	"fmt"
	"sync"
	"time"
)

type LogManager struct {
	fileLogger *FileLogger
	httpLogger *HTTPLogger
	logChan    chan LogEntry
	wg         sync.WaitGroup
}

func NewLogManager(filePath string) *LogManager {
	return &LogManager{
		fileLogger: NewFileLogger(filePath),
		httpLogger: &HTTPLogger{},
		logChan:    make(chan LogEntry, 100),
	}
}

func (m *LogManager) StartWorkers(count int) {
	for i := 0; i < count; i++ {
		m.wg.Add(1)
		go func(id int) {
			defer m.wg.Done()
			for entry := range m.logChan {
				fmt.Printf("[Worker-%d] Processing log...\n", id)
				m.fileLogger.Write(entry)
				m.httpLogger.Send(entry)
			}
		}(i + 1)
	}
}

func (m *LogManager) Log(message string) {
	entry := LogEntry{
		Timestamp: time.Now(),
		Message:   message,
	}
	m.logChan <- entry
}

func (m *LogManager) Stop() {
	close(m.logChan)
	m.wg.Wait()
	fmt.Println("All loggers shut down gracefully.")
}
