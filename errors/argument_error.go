package errors

import (
	"fmt"
	"io"
)

type ArgumentError struct {
	*StdError
	argument string
}

func NewArgumentError(argument string, msg string) *ArgumentError {
	return &ArgumentError{
		StdError: &StdError{
			msg:   msg,
			code:  "ArgumentError",
			stack: callers(),
		},
		argument: argument,
	}
}

func NewArgumentErrorf(argument string, msg string, args ...interface{}) *ArgumentError {
	return &ArgumentError{
		StdError: &StdError{
			msg:   fmt.Sprintf(msg, args...),
			code:  "ArgumentError",
			stack: callers(),
		},
		argument: argument,
	}
}

func (e *ArgumentError) Argument() string {
	return e.argument
}

func (e *ArgumentError) Is(target error) bool {
	if n, ok := target.(*ArgumentError); ok {
		return e.code == n.code
	}
	return false
}

func (e *ArgumentError) WithArgument(arg string) *ArgumentError {
	e.argument = arg
	return e
}

func (e *ArgumentError) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			io.WriteString(s, e.code+": "+e.msg+" "+e.argument)
			if e.cause != nil {
				s.Write([]byte{'\n'})
				fmt.Fprintf(s, "%+v", e.Cause())
			}

			e.stack.Format(s, verb)
			return
		}
		fallthrough
	case 's':
		io.WriteString(s, e.code+": "+e.msg+" "+e.argument)
	case 'q':
		fmt.Fprintf(s, "%q", e.Error())
	}
}
