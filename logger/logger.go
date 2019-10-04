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

var S struct {
	Level  int
	prefix string
}

func SetPrefix(prefix string) {
	S.prefix = prefix
}

func SetLevel(level int) {
	S.Level = level
}

func Info(args ...interface{}) {
	if S.Level >= INFO {
		Println(args...)
	}
}

func Infof(format string, args ...interface{}) {
	if S.Level >= INFO {
		Printf(format + "\n", args...)
	}
}

func Fatal(args ...interface{}) {
	if S.Level >= INFO {
		Println(args...)
	}
	os.Exit(1)
}

func Fatalf(format string, args ...interface{}) {
	if S.Level >= INFO {
		Printf(format + "\n", args...)
	}
	os.Exit(1)
}

func Debug(args ...interface{}) {
	if S.Level >= DEBUG {
		Println(args...)
	}
}

func Debugf(format string, args ...interface{}) {
	if S.Level >= DEBUG {
		Printf(format +"\n", args...)
	}
}

func Printf(format string, args ...interface{}) {
	if S.Level != NONE {
		fmt.Printf(S.prefix + format, args...)
	}
}

func Println(args ...interface{}) {
	if S.Level != NONE {
		fmt.Print(S.prefix)
		fmt.Println(args...)
	}
}

