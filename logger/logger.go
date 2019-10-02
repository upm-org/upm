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
	Level int
	Prefix string
}

var Log UPMLogger

func (l *UPMLogger) Info(format string, args ...interface{}) {
	if l.Level >= INFO {
		l.Println(args...)
	}
}

func (l *UPMLogger) Infof(format string, args ...interface{}) {
	if l.Level >= INFO {
		l.Printf(format + "\n", args...)
	}
}

func (l *UPMLogger) Fatal(format string, args ...interface{}) {
	if l.Level >= INFO {
		l.Println(args...)
	}
	os.Exit(1)
}

func (l *UPMLogger) Fatalf(format string, args ...interface{}) {
	if l.Level >= INFO {
		l.Printf(format + "\n", args...)
	}
	os.Exit(1)
}

func (l *UPMLogger) Debug(args ...interface{}) {
	if l.Level >= DEBUG {
		l.Println(args...)
	}
}

func (l *UPMLogger) Debugf(format string, args ...interface{}) {
	if l.Level >= DEBUG {
		l.Printf(format +"\n", args...)
	}
}

func (l *UPMLogger) Printf(format string, args ...interface{}) {
	fmt.Printf(l.Prefix + format, args...)
}

func (l *UPMLogger) Println(args ...interface{}) {
	fmt.Print(l.Prefix)
	fmt.Println(args...)
}

