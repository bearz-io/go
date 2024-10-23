package errors

import "fmt"

// AggregateError is a standard error for when multiple errors are aggregated
type AggregateError struct {
	*StdError
	errors []error
}

// NewAggregateError creates a new aggregate error with the given message.
// The code is set to "AggregateError".
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

// NewAggregateErrorf creates a new aggregate error with the given formatted message
// and the code is set to "AggregateError"
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

// Errors returns the errors in the aggregate error
func (e *AggregateError) Errors() []error {
	return e.errors
}

// Add adds an error to the aggregate error
func (e *AggregateError) Add(err error) {
	e.errors = append(e.errors, err)
}

// Is returns true if the error is the same as the target error
func (e *AggregateError) Is(target error) bool {
	if n, ok := target.(*AggregateError); ok {
		return e.code == n.code
	}
	return false
}
