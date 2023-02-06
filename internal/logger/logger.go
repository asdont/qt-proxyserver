package logger

import "fmt"

type Level string

const (
	Fatal Level = "FATAL"
	Error       = "ERROR"
	Info        = "INFO"
)

func Logger(level Level, msg string) string {
	return fmt.Sprintf("[%s] %s\n", level, msg)
}
