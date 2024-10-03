// Waku's logging
package log

import (
	"fmt"
)

// Log Level
type Level int

// Waku's Log Levels
//
// Arranged from most to least verbose
const (
	TRACE Level = iota
	DEBUG
	INFO
	WARNING
	ERROR
	QUIET

	debugPrefix = "\x1b[34m[DEBUG]\x1b[0m "
	infoPrefix  = "\x1b[32m[INFO]\x1b[0m  "
	warnPrefix  = "\x1b[33m[WARN]\x1b[0m  "
	errorPrefix = "\x1b[31m[ERROR]\x1b[0m "
)

// Waku's current log level
var logLevel Level = WARNING

func SetLevel(l Level) error {
	switch l {
	case TRACE, DEBUG, INFO, WARNING, ERROR, QUIET:
		logLevel = l
		Debugf("set log level to %d\n", l)
		return nil
	}
	return fmt.Errorf("invalid log level: %d", l)
}

func GetLevel() Level {
	return logLevel
}
