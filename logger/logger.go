package logger

/*
	Simple logger
*/

import (
	"fmt"
	"os"
)

const (
	NONE = 0
	INFO = 1
	DEBUG = 2
)

type UPMLogger struct {
	lvl int
	prefix string
}

var Log UPMLogger

func (l *UPMLogger) Init(Lvl int) {
	Log.lvl = Lvl
}

func (l *UPMLogger) SetPrefix(prefix string) {
	l.prefix = prefix
}

func (l *UPMLogger) Info(format string, args ...interface{}) {
	if l.lvl >= INFO {
		l.Printf(format + "\n", args...)
	}
}

func (l *UPMLogger) Fatal(prefix string, args ...interface{}) {
	l.Error("FATAL: " + prefix, args...)
}

func (l *UPMLogger) Error(format string, args ...interface{}) {
	if l.lvl >= INFO {
		l.Printf(format + "\n", args...)
	}
	os.Exit(1)
}

func (l *UPMLogger) Debug(format string, args ...interface{}) {
	if l.lvl >= DEBUG {
		l.Printf(format +"\n", args...)
	}
}

func (l *UPMLogger) Printf(format string, args ...interface{}) {
	fmt.Printf(l.prefix + format, args...)
}

