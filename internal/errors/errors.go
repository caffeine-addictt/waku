package errors

import (
	"fmt"
	"strings"
)

// WakuError is an error that should be returned in
// the CLI for logging to format messages properly
type WakuError struct {
	msg      string
	metadata []string
}

func (e *WakuError) WithMeta(key string, val string) *WakuError {
	return e.WithMetaf(key, "%s", val)
}

func (e *WakuError) WithMetaf(key string, format string, a ...any) *WakuError {
	e.metadata = append(e.metadata, fmt.Sprintf("%s=%s", key, fmt.Sprintf(format, a...)))
	return e
}

func (e *WakuError) Error() string {
	if len(e.metadata) == 0 {
		return e.msg
	}

	return fmt.Sprintf("[%s] %s", strings.Join(e.metadata, ", "), e.msg)
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
