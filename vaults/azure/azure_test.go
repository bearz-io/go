package azure_test

import (
    "testing"

    "github.com/bearz-io/go/vaults/azure"
    "github.com/stretchr/testify/assert"
)

func TestAzure(t *testing.T) {
    assert.Equal(t, azure.TEST, "TEST")
}