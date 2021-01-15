package cdep

import "fmt"

// Error logs a message with error level.
func Error(format string, args ...interface{}) {
	fmt.Printf("ERROR: %s\n", fmt.Sprintf(format, args...))
}

// Info logs a message with information level.
func Info(format string, args ...interface{}) {
	fmt.Printf("INFO: %s\n", fmt.Sprintf(format, args...))
}
