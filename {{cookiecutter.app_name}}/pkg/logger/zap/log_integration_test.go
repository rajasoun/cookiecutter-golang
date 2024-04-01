//go:build integration
// +build integration

package log_test

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/pkg/logger/log"
	"go.uber.org/zap"
)

func TestLogFunctions(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		logFunc  func(msg string, fields ...zap.Field)
		expected string
	}{
		{
			name:     "Info",
			logFunc:  log.Info,
			expected: "INFO",
		},
		{
			name:     "Warn",
			logFunc:  log.Warn,
			expected: "WARN",
		},
		{
			name:     "Error",
			logFunc:  log.Error,
			expected: "ERROR",
		},
		{
			name:     "Debug",
			logFunc:  log.Debug,
			expected: "DEBUG",
		},
	}

	// Iterate over the test cases
	for _, tc := range tests {
		// Create a buffer to capture log output
		var buf bytes.Buffer

		// Update the logger output to the buffer
		log.SetOutput(&buf)

		// Call the log function with a message
		tc.logFunc("Hello, world!")

		// Check that the buffer contains the log message and expected log level
		assert.Contains(t, buf.String(), "Hello, world!")
		assert.Contains(t, buf.String(), tc.expected)
	}
}

func TestLogger(t *testing.T) {
	// Capture the original stdout output
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Create a new logger with stdout output
	logger := log.NewLogger(os.Stdout)

	// Set the logger
	log.SetLogger(logger)

	// Log a message with Info level severity
	log.Info("Info message")

	// Read the stdout output
	w.Close()
	out, _ := io.ReadAll(r)

	// Ensure that the message was logged to stdout
	assert.Contains(t, string(out), "Info message")

	// Capture the new stdout output
	r, w, _ = os.Pipe()
	os.Stdout = w

	// Log a message with Debug level severity
	log.Debug("Debug message")

	// Read the stdout output
	w.Close()
	out, _ = io.ReadAll(r)

	// Ensure that the message was not logged to stdout
	assert.NotContains(t, string(out), "Debug message")

	// Reset the stdout output
	os.Stdout = old
}
