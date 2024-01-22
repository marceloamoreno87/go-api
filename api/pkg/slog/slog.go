package slog

import (
	"log/slog"
	"os"
)

type LoggerInterface interface {
	Debug(message string)
	Info(message string)
	Warning(message string)
	Error(message string)
}

type Logger struct {
	slog *slog.Logger
}

func NewLogger() *Logger {
	newlog := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(newlog)
	return &Logger{
		slog: newlog,
	}
}

func (l *Logger) Debug(message string) {
	slog.Debug(message)
}

func (l *Logger) Info(message string) {
	slog.Info(message)
}

func (l *Logger) Warning(message string) {
	slog.Warn(message)
}

func (l *Logger) Error(message string) {
	slog.Error(message)
}
