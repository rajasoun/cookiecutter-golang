// Package log provides a simple logger using Uber Zap logging.
package log

import (
	"io"
	"os"
	"path/filepath"

	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/pkg/testutils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

// init sets the logger output to os.Stdout by default.
func init() {
	core := getCore(os.Stdout)
	logger = zap.New(core)
}

// getCore creates a Zap core with a console encoder and the specified output writer.
func getCore(w io.Writer) zapcore.Core {
	consoleCore := getConsoleCore(w)
	// check if invoked from go test command - if so skip file logging
	if testutils.IsRunningAsGoTest() {
		return consoleCore
	}

	fileWriter := getFileWriter()
	fileCore := getFileCore(fileWriter)

	// Use zapcore.NewTee to combine the console and file cores into a single core that logs to both.
	core := zapcore.NewTee(consoleCore, fileCore)
	// Combine the console and file loggers into a single logger
	return core
}

func getFileWriter() io.Writer {
	logFilePath := filepath.Join(".", ".reports/logs")
	fileWriter, err := GetFileWriter(logFilePath)
	if err != nil {
		logger.Error("failed to get file writer", zap.String("error", err.Error()))
	}
	return fileWriter
}

// Info logs a message with Info level severity.
func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

// Warn logs a message with Warn level severity.
func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

// Error logs a message with Error level severity.
func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

// Debug logs a message with Debug level severity.
func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

// Fatal logs a message with Fatal level severity and exits the process.
func Fatal(msg string, fields ...zap.Field) {
	logger.Fatal(msg, fields...)
}

// NewLogger creates a new logger with a console encoder and the specified output writer.
func NewLogger(w io.Writer) *zap.Logger {
	core := getConsoleCore(w)
	return zap.New(core)
}

// SetOutput updates the logger output to the specified writer.
func SetOutput(w io.Writer) {
	core := getCore(w)
	logger = zap.New(core)
}

// SetLogger sets the logger for the log package.
func SetLogger(l *zap.Logger) {
	logger = l
}
