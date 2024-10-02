package log

import "fmt"

func Debugf(format string, a ...any) {
	if logLevel <= DEBUG {
		fmt.Printf(debugPrefix+format, a...)
	}
}

func Debugln(a ...any) {
	if logLevel <= DEBUG {
		fmt.Println(append([]any{debugPrefix}, a...)...)
	}
}

func Debug(v ...any) {
	if logLevel <= DEBUG {
		fmt.Print(append([]any{debugPrefix}, v...)...)
	}
}
