package paths_test

import (
    "testing"

    "github.com/bearz-io/go/os/paths"
    "github.com/stretchr/testify/assert"
)

func TestPaths(t *testing.T) {
    assert.Equal(t, paths.TEST, "TEST")
}