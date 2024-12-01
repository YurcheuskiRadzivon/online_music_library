package logger

import (
	"log"
)

type Logger struct {
	InfoLogger  *log.Logger
	DebugLogger *log.Logger
	ErrorLogger *log.Logger
}

func NewLogger() *Logger {
	return &Logger{
		InfoLogger:  log.New(log.Writer(), "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		DebugLogger: log.New(log.Writer(), "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile),
		ErrorLogger: log.New(log.Writer(), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}
