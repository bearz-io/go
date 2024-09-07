package strings_test

import (
    "testing"

    "github.com/bearz-io/go/strings"
    "github.com/stretchr/testify/assert"
)

func TestStrings(t *testing.T) {
    assert.Equal(t, strings.TEST, "TEST")
}