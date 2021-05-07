package logger

import "go.uber.org/zap"

// Logger is a trust to define logger object
type Logger struct {
	Log *zap.Logger
}

// New is a function to initialize zap logger
func New() *Logger {
	// setup logger
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	return &Logger{
		Log: logger,
	}
}

// Sync ...
func (l *Logger) Sync() {
	l.Log.Sync()
}
