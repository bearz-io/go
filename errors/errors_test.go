package errors_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/bearz-io/errors"
	"github.com/stretchr/testify/assert"
)

func TestErrors(t *testing.T) {
	e := errors.New("test")
	assert.Equal(t, e.Error(), "test")
}

func TestStdErrors(t *testing.T) {
	e := errors.NewStdError("test", "TestError")
	assert.Equal(t, "test", e.Error())
	assert.Equal(t, "TestError", e.Code())
	assert.Equal(t, e.Is(errors.NewStdError("test test", "TestError")), true)
	assert.Equal(t, e.Is(errors.NewStdError("test", "TestError2")), false)
	msg := fmt.Sprintf("%s", e)
	assert.Equal(t, "test", msg)
	msg = fmt.Sprintf("%+v", e)
	assert.True(t, strings.HasPrefix(msg, "TestError: test"))
}

func TestArgumentErrors(t *testing.T) {
	e := errors.NewArgumentError("test", "test")
	assert.Equal(t, e.Error(), "test")
	assert.Equal(t, e.Code(), "ArgumentError")
	assert.Equal(t, e.Is(errors.NewArgumentError("test", "test")), true)
	assert.Equal(t, e.Is(errors.NewArgumentError("test", "test2")), true)
	assert.Equal(t, e.Argument(), "test")
}
