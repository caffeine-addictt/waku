package log

import "fmt"

func Warnf(format string, a ...any) {
	if logLevel <= WARNING {
		fmt.Printf(warnPrefix+format, a...)
	}
}

func Warnln(a ...any) {
	if logLevel <= WARNING {
		fmt.Print(warnPrefix)
		fmt.Println(a...)
	}
}

func Warn(v ...any) {
	if logLevel <= WARNING {
		fmt.Print(warnPrefix)
		fmt.Print(v...)
	}
}
