package log

import (
	"fmt"
)

func Warnf(format string, a ...any) {
	if logLevel <= WARNING {
		fmt.Fprintf(Stderr, warnPrefix+format, a...)
	}
}

func Warnln(a ...any) {
	if logLevel <= WARNING {
		fmt.Fprint(Stderr, warnPrefix)
		fmt.Fprintln(Stderr, a...)
	}
}

func Warn(v ...any) {
	if logLevel <= WARNING {
		fmt.Fprint(Stderr, warnPrefix)
		fmt.Fprint(Stderr, v...)
	}
}
