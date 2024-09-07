package errors

import "fmt"

type AggregateError struct {
	*StdError
	errors []error
}

func NewAggregateError(errs []error, msg string) *AggregateError {
	return &AggregateError{
		StdError: &StdError{
			msg:   msg,
			code:  "AggregateError",
			stack: callers(),
		},
		errors: errs,
	}
}

func NewAggregateErrorf(errs []error, msg string, args ...interface{}) *AggregateError {
	return &AggregateError{
		StdError: &StdError{
			msg:   fmt.Sprintf(msg, args...),
			code:  "AggregateError",
			stack: callers(),
		},
		errors: errs,
	}
}

func (e *AggregateError) Errors() []error {
	return e.errors
}

func (e *AggregateError) Add(err error) {
	e.errors = append(e.errors, err)
}

func (e *AggregateError) Is(target error) bool {
	if n, ok := target.(*AggregateError); ok {
		return e.code == n.code
	}
	return false
}
