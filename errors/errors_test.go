package errors_test

import (
	"fmt"
	"testing"

	"github.com/bearz-io/go/errors"
	"github.com/stretchr/testify/assert"
)

func TestErrors(t *testing.T) {

	e := errors.New("test")
	assert.Equal(t, e.Error(), "test")
	fmt.Printf("%+v", e)
}
