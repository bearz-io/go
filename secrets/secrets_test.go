package secrets_test

import (
    "testing"

    "github.com/bearz-io/go/secrets"
    "github.com/stretchr/testify/assert"
)

func TestSecrets(t *testing.T) {
    assert.Equal(t, secrets.TEST, "TEST")
}