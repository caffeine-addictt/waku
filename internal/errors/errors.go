package errors

import "fmt"

// WakuError is an error that should be returned in
// the CLI for logging to format messages properly
type WakuError struct {
	msg string
}

// NewWakuError creates a new WakuError
func NewWakuErrorf(format string, v ...any) *WakuError {
	return &WakuError{msg: fmt.Sprintf(format, v...)}
}

func ToWakuError(err error) *WakuError {
	return &WakuError{msg: err.Error()}
}

func IsWakuError(err error) (*WakuError, bool) {
	w, ok := err.(*WakuError)
	return w, ok
}

func (e WakuError) Error() string {
	return e.msg
}
