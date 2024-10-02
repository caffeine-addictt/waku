package log

import "fmt"

func Infof(format string, a ...any) {
	if logLevel <= INFO {
		fmt.Printf(infoPrefix+format, a...)
	}
}

func Infoln(a ...any) {
	if logLevel <= INFO {
		fmt.Print(infoPrefix)
		fmt.Println(a...)
	}
}

func Info(v ...any) {
	if logLevel <= INFO {
		fmt.Print(infoPrefix)
		fmt.Print(v...)
	}
}
