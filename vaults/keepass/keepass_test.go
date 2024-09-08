package keepass_test

import (
    "testing"

    "github.com/bearz-io/go/vaults/keepass"
    "github.com/stretchr/testify/assert"
)

func TestKeepass(t *testing.T) {
    assert.Equal(t, keepass.TEST, "TEST")
}