package log

import (
	"fmt"
)

func Errorf(format string, a ...any) {
	if logLevel <= ERROR {
		fmt.Fprintf(Stderr, errorPrefix+format, a...)
	}
}

func Errorln(a ...any) {
	if logLevel <= ERROR {
		fmt.Fprint(Stderr, errorPrefix)
		fmt.Fprintln(Stderr, a...)
	}
}

func Error(v ...any) {
	if logLevel <= ERROR {
		fmt.Fprint(Stderr, errorPrefix)
		fmt.Fprint(Stderr, v...)
	}
}
