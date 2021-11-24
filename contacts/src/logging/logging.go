package logging

import (
	"log"
	"os"
	"strings"
)

// Init Logging
func Init() {
	logName := os.Getenv("log_filename")

	if strings.TrimSpace(logName) == "" {
		logName = "go-contacts.log"
	}

	file, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Print(err)
	}

	log.SetOutput(file)

	log.Print("Logging init started successfully")
}
