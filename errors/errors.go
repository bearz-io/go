package errors

import (
	"fmt"
	"runtime"
)

var (
	ErrNotFound        = NewResourceError("resource not found", "resource")
	ErrOsNotSupported  = NewResourceError("os not supported", runtime.GOOS)
	ErrInvalidOp       = NewStdError("invalid operation", "InvalidOperation")
	ErrNotImpelemented = NewResourceError("not implemented", "feature")
	ErrArgEmpty        = NewArgumentError("argument is empty", "unknown")
	ErrArgNil          = NewArgumentError("argument is nil", "unknown")
)

func New(msg string) error {
	return &StdError{
		msg:   msg,
		stack: callers(),
		code:  "Error",
	}
}

func Newf(msg string, args ...interface{}) error {
	return &StdError{
		msg:   fmt.Sprintf(msg, args...),
		stack: callers(),
		code:  "Error",
	}
}

func NewStdError(msg string, code string) *StdError {
	return &StdError{
		msg:   msg,
		code:  code,
		stack: callers(),
	}
}

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
		code:  "StdError",
		cause: err,
		stack: callers(),
	}
}

func Errorf(msg string, args ...interface{}) error {
	return &StdError{
		msg:   fmt.Sprintf(msg, args...),
		stack: callers(),
		code:  "StdError",
	}
}
