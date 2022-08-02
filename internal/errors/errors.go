package errors

import (
	"fmt"
	"runtime"
)

// Err --
type Err struct {
	HTTPCode int    // forward to client
	Message  string // forward to client
	Source   string // for logging
	Err      error  // for logging
}

// New init errors
func New(httpCode int, msg string) *Err {
	return &Err{
		HTTPCode: httpCode,
		Message:  msg,
	}
}

func (e *Err) Error() string {
	return e.Message
}

// Warp will capture original error and warp it with custom error
func (e *Err) Warp(errs ...error) error {
	if len(errs) > 0 {
		// only save first original error
		e.Err = errs[0]
	}

	// init error source
	if pc, file, line, ok := runtime.Caller(1); ok {
		details := runtime.FuncForPC(pc)

		// [main.main] /tmp/sandbox1258520468/prog.go:16
		e.Source = fmt.Sprintf("[%s] %s:%d", details.Name(), file, line)
	}

	return e
}

// SetMessage will overide existing error message
func (e *Err) SetMessage(msg ...string) *Err {
	if len(msg) > 0 {
		e.Message = msg[0]
	}
	return e
}
