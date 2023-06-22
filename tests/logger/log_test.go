package logger_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"smpt_server/logger"
	"strings"
	"testing"
)

func TestSetupLogger(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir, err := ioutil.TempDir("", "logger_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	// Set the current working directory to the temporary directory
	prevDir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	defer os.Chdir(prevDir)
	err = os.Chdir(tmpDir)
	if err != nil {
		t.Fatal(err)
	}

	// Call the function being tested
	log, err := logger.SetupLogger()
	if err != nil {
		t.Fatal(err)
	}

	// Verify that the log file exists
	logFilePath := filepath.Join(tmpDir, "app.log")
	_, err = os.Stat(logFilePath)
	if err != nil {
		t.Errorf("Log file does not exist: %s", logFilePath)
	}

	// Verify that the logger writes to the log file
	expectedLogMessage := "Test log message"
	log.Println(expectedLogMessage)

	// Read the log file
	data, err := ioutil.ReadFile(logFilePath)
	if err != nil {
		t.Fatal(err)
	}

	// Verify that the log message is written to the file
	actualLogMessage := string(data)
	if !strings.Contains(actualLogMessage, expectedLogMessage) {
		t.Errorf("Expected log message '%s' not found in '%s'", expectedLogMessage, actualLogMessage)
	}
}
