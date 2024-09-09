package hostfile_test

import (
    "testing"

    "github.com/bearz-io/go/os/hostfile"
    "github.com/stretchr/testify/assert"
)

func TestHostfile(t *testing.T) {
    assert.Equal(t, hostfile.TEST, "TEST")
}