package log_test

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
)

// MockLogger is a mock logger that implements the zap.Logger interface.
type MockLogger struct {
	mock.Mock
}

// Info logs a message with Info level severity.
func (m *MockLogger) Info(msg string, fields ...zap.Field) {
	m.Called(msg, fields)
}

// Warn logs a message with Warn level severity.
func (m *MockLogger) Warn(msg string, fields ...zap.Field) {
	m.Called(msg, fields)
}

// Error logs a message with Error level severity.
func (m *MockLogger) Error(msg string, fields ...zap.Field) {
	m.Called(msg, fields)
}

// Debug logs a message with Debug level severity.
func (m *MockLogger) Debug(msg string, fields ...zap.Field) {
	m.Called(msg, fields)
}

// Trace logs a message with Trace level severity.
func (m *MockLogger) Trace(msg string, fields ...zap.Field) {
	m.Called(msg, fields)
}

// Trace logs a message with Trace level severity.
func (m *MockLogger) Fatal(msg string, fields ...zap.Field) {
	m.Called(msg, fields)
}

func LogMe(logger *MockLogger) {
	// Log an info message
	logger.Info("Info message")

	// Log an error message
	logger.Warn("Warn message")

	// Log an error message
	logger.Error("Error message")

	// Log a debug message
	logger.Debug("Debug message")

	// Log a fatal message
	logger.Fatal("Fatal message")

}

func TestLogMe(t *testing.T) {
	t.Parallel()
	// Create a mock logger
	logger := &MockLogger{}

	// Expect the methods to be called with the expected arguments
	logger.On("Info", "Info message", mock.AnythingOfType("[]zapcore.Field"))
	logger.On("Warn", "Warn message", mock.AnythingOfType("[]zapcore.Field"))
	logger.On("Error", "Error message", mock.AnythingOfType("[]zapcore.Field"))
	logger.On("Debug", "Debug message", mock.AnythingOfType("[]zapcore.Field"))
	logger.On("Fatal", "Fatal message", mock.AnythingOfType("[]zapcore.Field"))

	// Call the function to be tested, passing in the mock logger
	LogMe(logger)

	// Assert that the expected methods were called
	logger.AssertExpectations(t)
}
