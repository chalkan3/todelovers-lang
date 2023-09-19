package logger

import (
	"fmt"
	"io"
	"os"
)

// Logger interface
type Logger interface {
	// Log a message
	Log(level int, message string)
}

// ConsoleLogger struct
type ConsoleLogger struct {
	// Output stream
	output io.Writer
}

// NewConsoleLogger function
func NewConsoleLogger(output io.Writer) *ConsoleLogger {
	return &ConsoleLogger{output: output}
}

// Log method
func (l *ConsoleLogger) Log(level int, message string) {
	// Format the message
	ss := [...]string{"INFO", "WARNING", "ERROR"}[level]
	formattedMessage := fmt.Sprintf("[%s]%s\n", ss, message)

	// Write the message to the output stream
	l.output.Write([]byte(formattedMessage))
}

// FileLogger struct
type FileLogger struct {
	// Output file
	file *os.File
}

// NewFileLogger function
func NewFileLogger(filename string) (*FileLogger, error) {
	// Open the file
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	// Create the logger
	return &FileLogger{file: file}, nil
}

// Log method
func (l *FileLogger) Log(level int, message string) {
	// Format the message
	ss := [...]string{"INFO", "WARNING", "ERROR"}[level]

	formattedMessage := fmt.Sprintf("[%s]%s\n", ss, message)

	// Write the message to the file
	l.file.WriteString(formattedMessage)
}
