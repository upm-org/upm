package logger

/*
	Simple logger
*/

import (
	"log"
	"os"
)

const (
	NONE = 0
	INFO = 1
	DEBUG = 2
)

type UPMLogger struct {
	Lvl int
}

func (l *UPMLogger) Info(format string, args ...interface{}) {
	if (l.Lvl >= INFO) {
		l.Printf(format, args...)
	}
}

func (l *UPMLogger) Error(format string, args ...interface{}) {
	if (l.Lvl >= INFO) {
		l.Printf(format, args...)
	}
	os.Exit(1)
}

func (l *UPMLogger) Debug(format string, args ...interface{}) {
	if (l.Lvl >= DEBUG) {
		l.Printf(format, args...)
	}
}

func (l *UPMLogger) Printf(format string, args ...interface{}) {
	log.Printf(format, args...)
}


