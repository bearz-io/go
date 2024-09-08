package sops_test

import (
    "testing"

    "github.com/bearz-io/go/vaults/sops"
    "github.com/stretchr/testify/assert"
)

func TestSops(t *testing.T) {
    assert.Equal(t, sops.TEST, "TEST")
}