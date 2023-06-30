package stackerr

import (
	"fmt"
	"io"
	"runtime"
)

type StackTraceError struct {
	err        error
	stackTrace string
}

func NewStackTraceError(err error, msg string) StackTraceError {
	stackBuf := make([]byte, 1024)
	n := runtime.Stack(stackBuf, false)
	stackTrace := string(stackBuf[:n])

	return StackTraceError{
		err:        fmt.Errorf("%s: %w", msg, err),
		stackTrace: stackTrace,
	}
}

func (e StackTraceError) Error() string {
	return e.err.Error()
}

func (e StackTraceError) Unwrap() error {
	return e.err
}

func (e StackTraceError) StackTrace() string {
	return e.stackTrace
}

func (e StackTraceError) PrintStackTrace(w io.Writer) {
	fmt.Fprintln(w, e.stackTrace)
}
