package log

import "fmt"

func Infof(format string, a ...any) {
	if logLevel <= INFO {
		fmt.Fprintf(Stdout, infoPrefix+format, a...)
	}
}

func Infoln(a ...any) {
	if logLevel <= INFO {
		fmt.Fprint(Stdout, infoPrefix)
		fmt.Fprintln(Stdout, a...)
	}
}

func Info(v ...any) {
	if logLevel <= INFO {
		fmt.Fprint(Stdout, infoPrefix)
		fmt.Fprint(Stdout, v...)
	}
}
