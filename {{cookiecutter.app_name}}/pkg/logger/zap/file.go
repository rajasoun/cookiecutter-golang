package log

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const logDir = "./.reports/logs"

// getFileCore creates a Zap core with a file encoder and the specified output writer.
func getFileCore(w io.Writer) zapcore.Core {
	// Configure the encoder
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.Lock(zapcore.AddSync(w)),
		zap.NewAtomicLevelAt(zap.InfoLevel),
	)

	return core
}

// GetFileWriter returns a file writer for the specified file path.
func GetFileWriter(filePath string) (io.Writer, error) {
	// Configure the file name and directory
	logFileName := filepath.Join(logDir, "log-"+time.Now().Format("2006-01-02")+".json")

	// Create the log directory if it doesn't exist
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("failed to create directory: %w", err)
		}
		// logDir already exists, ignore error
		Warn("logDir already exists, ignore error", zap.Error(err))
	}

	// Create the file writer
	file, err := os.OpenFile(filepath.Clean(logFileName), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}

	// Convert the *os.File to an io.Writer using os.Stdout as a reference.
	fileWriter := io.Writer(file)
	return fileWriter, nil
}
