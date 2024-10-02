package log

import "fmt"

func Printf(format string, a ...any) {
	if logLevel != QUIET {
		fmt.Printf(format, a...)
	}
}

func Println(a ...any) {
	if logLevel != QUIET {
		fmt.Println(a...)
	}
}

func Print(v ...any) {
	if logLevel != QUIET {
		fmt.Print(v...)
	}
}
