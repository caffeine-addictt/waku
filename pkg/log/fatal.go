package log

import (
	"fmt"
	"os"
	"runtime/debug"
)

func Fatalf(format string, a ...any) {
	Errorf(format, a...)
	if logLevel <= TRACE {
		fmt.Fprintln(Stderr, string(debug.Stack()))
	}
	os.Exit(1)
}

func Fatalln(a ...any) {
	Errorln(a...)
	if logLevel <= TRACE {
		fmt.Fprintln(Stderr, string(debug.Stack()))
	}
	os.Exit(1)
}

func Fatal(v ...any) {
	Error(v...)
	if logLevel <= TRACE {
		fmt.Fprintln(Stderr, string(debug.Stack()))
	}
	os.Exit(1)
}
