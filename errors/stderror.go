package errors

import (
	"fmt"
	"io"
)

type StdError struct {
	*stack
	msg   string
	cause error
	code  string
}

func (e *StdError) Error() string {
	if e.cause != nil {
		return e.msg + ": " + e.cause.Error()
	}

	return e.msg
}

func (e *StdError) Code() string {
	return e.code
}

func (e *StdError) Cause() error {
	return e.cause
}

func (e *StdError) Unwrap() error {
	return e.cause
}

func (e *StdError) Is(target error) bool {
	if n, ok := target.(*StdError); ok {
		return e.code == n.code
	}

	return false
}

func (e *StdError) WithCause(err error) *StdError {
	e.cause = err
	return e
}

func (e *StdError) WithMessage(msg string) *StdError {
	e.msg = msg
	return e
}

func (e *StdError) WithMessageF(msg string, args ...interface{}) *StdError {
	e.msg = fmt.Sprintf(msg, args...)
	return e
}

func (e *StdError) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			io.WriteString(s, e.code+": "+e.msg)
			if e.cause != nil {
				s.Write([]byte{'\n'})
				fmt.Fprintf(s, "%+v", e.Cause())
			}

			e.stack.Format(s, verb)
			return
		}
		fallthrough
	case 's':
		io.WriteString(s, e.Error())
	case 'q':
		fmt.Fprintf(s, "%q", e.Error())
	}
}
