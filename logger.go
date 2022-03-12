package recurly

import (
	"fmt"
	"log"
	"os"
)

const (
	LevelDebug Level = 0
	LevelInfo  Level = 1
	LevelWarn  Level = 2
	LevelError Level = 3
)

// Level log level
type Level int

// Logger creates a new logger with variable level
type Logger struct {
	Level
	StdOut *log.Logger
	StdErr *log.Logger
}

func NewLogger(logLevel Level) *Logger {
	return &Logger{
		Level:  logLevel,
		StdOut: log.New(os.Stdout, "Recurly ", log.LstdFlags|log.Lmicroseconds),
		StdErr: log.New(os.Stderr, "Recurly ", log.LstdFlags|log.Lmicroseconds),
	}
}

// IsLevel returns true if the logger would print with the given level
func (log *Logger) IsLevel(level Level) bool {
	return log.Level <= level
}

func (log *Logger) Debug(v ...interface{}) {
	if log.IsLevel(LevelDebug) {
		log.stdOut(fmt.Sprint(v...))
	}
}

func (log *Logger) Debugf(format string, v ...interface{}) {
	if log.IsLevel(LevelDebug) {
		log.stdOut(fmt.Sprintf(format, v...))
	}
}

func (log *Logger) Info(v ...interface{}) {
	if log.IsLevel(LevelInfo) {
		log.stdOut(fmt.Sprint(v...))
	}
}

func (log *Logger) Infof(format string, v ...interface{}) {
	if log.IsLevel(LevelInfo) {
		log.stdOut(fmt.Sprintf(format, v...))
	}
}

func (log *Logger) Warn(v ...interface{}) {
	if log.IsLevel(LevelWarn) {
		log.stdOut(fmt.Sprint(v...))
	}
}

func (log *Logger) Warnf(format string, v ...interface{}) {
	if log.IsLevel(LevelWarn) {
		log.stdOut(fmt.Sprintf(format, v...))
	}
}

func (log *Logger) Error(v ...interface{}) {
	if log.IsLevel(LevelError) {
		log.stdOut(fmt.Sprint(v...))
	}
}

func (log *Logger) Errorf(format string, v ...interface{}) {
	if log.IsLevel(LevelError) {
		log.stdOut(fmt.Sprintf(format, v...))
	}
}

func (log *Logger) stdOut(v ...interface{}) {
	if log.StdOut != nil {
		log.StdOut.Print(v...)
	}
}

//lint:ignore U1000 TODO: should `Error` call stdErr instead of stdOut?
func (log *Logger) stdErr(v ...interface{}) {
	if log.StdErr != nil {
		log.StdErr.Print(v...)
	}
}
