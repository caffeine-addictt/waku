package log

import "fmt"

func Debugf(format string, a ...any) {
	if logLevel <= DEBUG {
		fmt.Fprintf(Stdout, debugPrefix+format, a...)
	}
}

func Debugln(a ...any) {
	if logLevel <= DEBUG {
		fmt.Fprint(Stdout, debugPrefix)
		fmt.Fprintln(Stdout, a...)
	}
}

func Debug(v ...any) {
	if logLevel <= DEBUG {
		fmt.Fprint(Stdout, debugPrefix)
		fmt.Fprint(Stdout, v...)
	}
}
