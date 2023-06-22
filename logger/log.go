package logger

import (
	"log"
	"os"
)

func SetupLogger() (*log.Logger, error) {
	// Open the log file in append mode
	file, err := os.OpenFile("app.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}

	// Create a new logger that writes to the log file
	return log.New(file, "", log.LstdFlags), nil
}
