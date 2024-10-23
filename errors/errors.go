package errors

import (
	"fmt"
	"runtime"
)

var (
	// ErrNotFound is a standard error for when a resource is not found
	ErrNotFound = NewResourceError("resource not found", "resource")

	// ErrOsNotSupported is a standard error for when the operating system is not supported
	ErrOsNotSupported = NewResourceError("os not supported", runtime.GOOS)

	// ErrInvalidOp is a standard error for when an invalid operation is attempted
	ErrInvalidOp = NewStdError("invalid operation", "InvalidOperation")

	// ErrNotSupported is a standard error for when a feature is not supported
	ErrNotSupported = NewResourceError("not supported", "resource")

	// ErrNotImplemented is a standard error for when a feature is not implemented
	ErrNotImplemented = NewResourceError("not implemented", "feature")

	// ErrArgEmpty is a standard error for when an argument is empty
	ErrArgEmpty = NewArgumentError("argument is empty", "unknown")

	// ErrArgNil is a standard error for when an argument is nil
	ErrArgNil = NewArgumentError("argument is nil", "unknown")

	// ErrAccessDenied is a standard error for when access is denied
	ErrAccessDenied = NewResourceError("access denied", "resource")
)

// New creates a new standard error with the given message.
// The code is set to "Error"
//
//	err := errors.New("program failed")
//	fmt.Printf("%+v", err)
func New(msg string) error {
	return &StdError{
		msg:   msg,
		stack: callers(),
		code:  "Error",
	}
}

// Newf creates a new standard error with the given formatted message
// and the code is set to "Error"
//
//	err := errors.Newf("program failed with %s", "error")
//	fmt.Printf("%+v", err)
func Newf(msg string, args ...interface{}) error {
	return &StdError{
		msg:   fmt.Sprintf(msg, args...),
		stack: callers(),
		code:  "Error",
	}
}

// NewStdError creates a new standard error with the given message and code
//
//	  var (
//	    MyCustomError = errors.NewStdError("program failed", "MyCustomError")
//	  )
//
//	  func main() {
//	    args := os.Args[1:]
//		if len(args) == 0 {
//		  fmt.Printf("%+v", MyCustomError)
//		}
//	  }
func NewStdError(msg string, code string) *StdError {
	return &StdError{
		msg:   msg,
		code:  code,
		stack: callers(),
	}
}

// Join joins multiple errors into a single error. The
// underlying error is an AggregateError which includes
// the code and stack from where Join was called.
func Join(errs ...error) error {

	if len(errs) == 0 {
		return nil
	}

	if len(errs) == 1 {
		return errs[0]
	}

	return &AggregateError{
		StdError: &StdError{
			msg:   "multiple errors",
			code:  "AggregateError",
			stack: callers(),
		},
		errors: errs,
	}
}

// Wrap wraps an error with a message. If the error is nil, it returns nil.
// If the error is already a StdError, it sets the message and returns the error.
// Otherwise, it returns a new StdError with the message and the code set to "Error".
func Wrap(err error, msg string) error {
	if err == nil {
		return nil
	}

	if n, ok := err.(*StdError); ok {
		n.msg = msg
		return err
	}

	return &StdError{
		msg:   msg,
		code:  "Error",
		cause: err,
		stack: callers(),
	}
}

// Errorf creates a new standard error with the given formatted message
// and the code is set to "Error"
func Errorf(msg string, args ...interface{}) error {
	return &StdError{
		msg:   fmt.Sprintf(msg, args...),
		stack: callers(),
		code:  "Error",
	}
}
