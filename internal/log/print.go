package log

import "fmt"

func Printf(format string, a ...any) {
	if logLevel != QUIET {
		fmt.Fprintf(Stdout, format, a...)
	}
}

func Println(a ...any) {
	if logLevel != QUIET {
		fmt.Fprintln(Stdout, a...)
	}
}

func Print(v ...any) {
	if logLevel != QUIET {
		fmt.Fprint(Stdout, v...)
	}
}
