package logger

import (
	"log"
)

type Logger struct{}

func (l *Logger) Error(s string, err error) {
	log.Printf("[ERROR] %s: %v", s, err) // Логируем ошибку вместо паники
}

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) Info(format string, v ...interface{}) {
	log.Printf("[INFO] "+format, v...)
}

func (l *Logger) Fatal(format string, v ...interface{}) {
	log.Fatalf("[FATAL] "+format, v...)
}
