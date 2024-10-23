package errors

import (
	"fmt"
	"io"
)

// ArgumentError is a standard error for when an argument is invalid
type ArgumentError struct {
	*StdError
	argument string
}

// NewArgumentError creates a new argument error with the given message and code
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

// NewArgumentErrorf creates a new argument error with the given formatted message
// and the code is set to "ArgumentError"
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

// Argument returns the argument of the error
func (e *ArgumentError) Argument() string {
	return e.argument
}

// Is returns true if the error is the same as the target error
func (e *ArgumentError) Is(target error) bool {
	if n, ok := target.(*ArgumentError); ok {
		return e.code == n.code
	}
	return false
}

// Sets the argument of the error
func (e *ArgumentError) WithArgument(arg string) *ArgumentError {
	e.argument = arg
	return e
}

// Formats the error. The verb is used to determine how to format the error
// '%+v' will print the error code, message, cause, and stacktrace
// '%v' will print the error code, message, and stacktrace
// '%s' will print the error message
// '%q' will print the error message using fmt.Sprintf("%q", e.Error())
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
