package log

import (
	"go.uber.org/zap"
)

// Created 20/11/2021
// Updated
type Logger interface {
	GetInstance() interface{}
	// Debug(...interface{})
	// Error(...interface{})
	// Fatal(...interface{})
	// Info(...interface{})
	// Panic(...interface{})
	// Warn(...interface{})
}

type SugaredLogger struct {
	Logger *zap.SugaredLogger
}

type ZapLogger struct {
	Logger *zap.Logger
}

func (log *ZapLogger) GetInstance() interface{} {
	return log.Logger
}

func NewSugarLogger() Logger {
	l, _ := zap.NewProduction()

	return &SugaredLogger{l.Sugar()}
}

func NewZapLogger() Logger {
	l, _ := zap.NewProduction()

	return &ZapLogger{l}
}

func (log *SugaredLogger) GetInstance() interface{} {
	return log.Logger
}

// type NewZapLogger()
