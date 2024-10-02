package log

import "fmt"

func Warnf(format string, a ...any) {
	if logLevel <= WARNING {
		fmt.Printf(warnPrefix+format, a...)
	}
}

func Warnln(a ...any) {
	if logLevel <= WARNING {
		fmt.Println(append([]any{warnPrefix}, a...)...)
	}
}

func Warn(v ...any) {
	if logLevel <= WARNING {
		fmt.Print(append([]any{warnPrefix}, v...)...)
	}
}
