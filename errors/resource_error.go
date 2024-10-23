package errors

import "fmt"

type ResourceError struct {
	*StdError
	resource string
}

func NewResourceError(resource string, msg string) *ResourceError {
	return &ResourceError{
		StdError: &StdError{
			msg:   msg,
			code:  "ResourceError",
			stack: callers(),
		},
		resource: resource,
	}
}

func NewResourceErrorf(resource string, msg string, args ...interface{}) *ResourceError {
	return &ResourceError{
		StdError: &StdError{
			msg:   fmt.Sprintf(msg, args...),
			code:  "ResourceError",
			stack: callers(),
		},
		resource: resource,
	}
}

func (e *ResourceError) Resource() string {
	return e.resource
}

func (e *ResourceError) Is(target error) bool {
	if n, ok := target.(*ResourceError); ok {
		return e.code == n.code
	}
	return false
}

func (e *ResourceError) WithResource(resource string) *ResourceError {
	e.resource = resource
	return e
}
