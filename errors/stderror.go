package errors

import (
	"fmt"
	"io"
)

// The standard error which supports code, cause, and stacktrace.
//
// The code is set to "Error" by default.  Code is useful for distinguishing
// between different types of errors and for the Is() function or the web
// where the code is used to determine the error type sent to the client.
//
// The cause is set to nil by default.  The cause is the underlying error
// that caused the error.  The cause is used to determine the root cause of
// the error and for the WithCause() function.
//
// The stack is set to the callers() function by default.  The stack is the
// stacktrace of the error.  The stack is used to determine the stacktrace
// of the error and for the WithStack() function.
//
// The StdError needs to be created with the NewStdError() function.
type StdError struct {
	*stack
	msg   string
	cause error
	code  string
}

// Error returns the error message
func (e *StdError) Error() string {
	return e.msg
}

// Code returns the error code
func (e *StdError) Code() string {
	return e.code
}

// Cause returns the underlying cause of the error, if one
// exists.
func (e *StdError) Cause() error {
	return e.cause
}

// Unwrap returns the underlying cause of the error, if one
// exists.
func (e *StdError) Unwrap() error {
	return e.cause
}

// Is returns true if the error is the same as the target error
func (e *StdError) Is(target error) bool {
	if n, ok := target.(*StdError); ok {
		return e.code == n.code
	}

	return false
}

// Sets the stacktrace
func (e *StdError) SetStack(stacktrace *[]uintptr) *StdError {
	if stacktrace == nil {
		return e
	} else {
		var st stack = *stacktrace
		e.stack = &st
		return e
	}
}

// Sets the error code
func (e *StdError) SetCode(code string) *StdError {
	e.code = code
	return e
}

// Sets the underlying cause of the error
func (e *StdError) SetCause(cause error) *StdError {
	e.cause = cause
	return e
}

// Creates a new StdError that clones the current StdError and sets the underlying cause of the error
// The code is set to "Error"
func (e *StdError) WithCause(err error) *StdError {
	e2 := &StdError{
		msg:   e.msg,
		code:  e.code,
		cause: err,
		stack: e.stack,
	}
	return e2
}

// Creates a new StdError that clones the current StdError and sets the error message
// The code is set to "Error"
func (e *StdError) WithMessage(msg string) *StdError {
	e2 := &StdError{
		msg:   msg,
		code:  e.code,
		cause: e.cause,
		stack: e.stack,
	}
	return e2
}

// Creates a new StdError that clones the current StdError and sets
// the error message with a formatted string.
//
//	var MyCustomError = errors.NewStdError("program failed", "MyCustomError")
//
//	// ... in a function
//	err := MyCustomError.WithMessageF("program failed with %s", "bad input")
func (e *StdError) WithMessageF(msg string, args ...interface{}) *StdError {
	msg2 := fmt.Sprintf(msg, args...)
	e2 := &StdError{
		msg:   msg2,
		code:  e.code,
		cause: e.cause,
		stack: e.stack,
	}
	return e2
}

// Creates a new StdError that clones the current StdError and sets the stacktrace.
// This is useful with using custom errors stored as a variable.
//
//	var MyCustomError = errors.NewStdError("program failed", "MyCustomError")
//	// ... in a function
//	err := MyCustomError.WithStack()
func (e *StdError) WithStack() *StdError {
	return &StdError{
		msg:   e.msg,
		code:  e.code,
		cause: e.cause,
		stack: callers(),
	}
}

// Updates the stacktrace when called from a function that
// has been called by the WithStack() function
func (e *StdError) UpdateStack() *StdError {
	e.stack = callers()
	return e
}

// Formats the error. The verb is used to determine how to format the error
// '%+v' will print the error code, message, cause, and stacktrace
// '%v' will print the error code, message, and stacktrace
// '%s' will print the error message
// '%q' will print the error message using fmt.Sprintf("%q", e.Error())
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
		} else {
			io.WriteString(s, e.code)
			io.WriteString(s, ": ")
			io.WriteString(s, e.Error())
			e.stack.Format(s, verb)
		}

	case 's':
		io.WriteString(s, e.msg)
	case 'q':
		fmt.Fprintf(s, "%q", e.Error())
	}
}
