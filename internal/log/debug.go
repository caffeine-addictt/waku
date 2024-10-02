package log

import "fmt"

func Debugf(format string, a ...any) {
	if logLevel <= DEBUG {
		fmt.Printf(debugPrefix+format, a...)
	}
}

func Debugln(a ...any) {
	if logLevel <= DEBUG {
		fmt.Print(debugPrefix)
		fmt.Println(a...)
	}
}

func Debug(v ...any) {
	if logLevel <= DEBUG {
		fmt.Print(debugPrefix)
		fmt.Print(v...)
	}
}
