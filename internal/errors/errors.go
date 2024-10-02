package errors

import "fmt"

// WakuError is an error that should be returned in
// the CLI for logging to format messages properly
type WakuError struct {
	msg string
}

// NewWakuError creates a new WakuError
func NewWakuErrorf(format string, v ...any) WakuError {
	return WakuError{msg: fmt.Sprintf(format, v...)}
}

func (e WakuError) Error() string {
	return e.msg
}
