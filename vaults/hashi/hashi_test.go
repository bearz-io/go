package hashi_test

import (
    "testing"

    "github.com/bearz-io/go/vaults/hashi"
    "github.com/stretchr/testify/assert"
)

func TestHashi(t *testing.T) {
    assert.Equal(t, hashi.TEST, "TEST")
}