package options

import (
	"log"
)

const (
	debugPrefix   = "\x1b[34m[DEBUG]\x1b[0m "
	verbosePrefix = "\x1b[32m[INFO]\x1b[0m "
)

func Debugf(format string, a ...any) {
	if GlobalOpts.Debug {
		log.Printf(debugPrefix+format, a...)
	}
}

func Debugln(a ...any) {
	if GlobalOpts.Debug {
		log.Println(append([]any{debugPrefix}, a...)...)
	}
}

func Debug(v ...any) {
	if GlobalOpts.Debug {
		log.Print(append([]any{debugPrefix}, v...)...)
	}
}

func Infof(format string, a ...any) {
	if GlobalOpts.DebugOrVerbose() {
		log.Printf(verbosePrefix+format, a...)
	}
}

func Infoln(a ...any) {
	if GlobalOpts.DebugOrVerbose() {
		log.Println(append([]any{verbosePrefix}, a...)...)
	}
}

func Info(v ...any) {
	if GlobalOpts.DebugOrVerbose() {
		log.Print(append([]any{verbosePrefix}, v...)...)
	}
}
