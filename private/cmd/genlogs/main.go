package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

const (
	numLogEntries      = 10000
	maxIntervalSeconds = 18
	startTime          = "2024-01-11 00:00:00"
)

func main() {
	start, err := time.Parse("2006-01-02 15:04:05", startTime)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create(start.Format("2006-01-02") + ".log")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for i := 0; i < numLogEntries; i++ {
		logEntry := generateLogEntry(start)
		fmt.Fprintln(file, logEntry)

		// Increment the time for the next log entry
		interval := time.Duration(rand.Intn(maxIntervalSeconds)) * time.Second
		start = start.Add(interval)
	}

	fmt.Println("Log file created successfully.")
}

func generateLogEntry(timestamp time.Time) string {
	messages := []string{
		"Server shutdown complete.",
		"Server rebooting.",
		"Server reboot complete. System ready.",
		"Database connection established successfully.",
		"Server startup complete. System ready.",
		"Server not connected to Network. Check network connection.",
		"Network connection re-established.",
		"Security scan initiated.",
		"Security scan completed. No threats found.",
	}

	levels := []string{"INFO", "WARNING", "ERROR", "ALERT"}

	message := messages[rand.Intn(len(messages))]
	level := levels[rand.Intn(len(levels))]

	return fmt.Sprintf("%s %s: %s", timestamp.Format("2006-01-02 15:04:05"), level, message)
}
