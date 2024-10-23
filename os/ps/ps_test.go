package ps_test

import (
    "testing"

    "github.com/bearz-io/go/os/ps"
    "github.com/stretchr/testify/assert"
)

func TestPs(t *testing.T) {
    assert.Equal(t, ps.TEST, "TEST")
}