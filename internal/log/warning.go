package log

import (
	"fmt"
	"os"
)

func Warnf(format string, a ...any) {
	if logLevel <= WARNING {
		fmt.Fprintf(os.Stderr, warnPrefix+format, a...)
	}
}

func Warnln(a ...any) {
	if logLevel <= WARNING {
		fmt.Fprint(os.Stderr, warnPrefix)
		fmt.Fprintln(os.Stderr, a...)
	}
}

func Warn(v ...any) {
	if logLevel <= WARNING {
		fmt.Fprint(os.Stderr, warnPrefix)
		fmt.Fprint(os.Stderr, v...)
	}
}
