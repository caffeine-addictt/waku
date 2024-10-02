package log

import (
	"fmt"
	"os"
)

func Errorf(format string, a ...any) {
	if logLevel <= ERROR {
		fmt.Fprintf(os.Stderr, errorPrefix+format, a...)
	}
}

func Errorln(a ...any) {
	if logLevel <= ERROR {
		fmt.Fprint(os.Stderr, errorPrefix)
		fmt.Fprintln(os.Stderr, a...)
	}
}

func Error(v ...any) {
	if logLevel <= ERROR {
		fmt.Fprint(os.Stderr, errorPrefix)
		fmt.Fprint(os.Stderr, v...)
	}
}
