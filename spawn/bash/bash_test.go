package bash_test

import (
	"strings"
	"testing"

	"github.com/bearz-io/go/spawn/bash"
	"github.com/stretchr/testify/assert"
)

func TestBash(t *testing.T) {
	hasBash := bash.Which() != nil
	if !hasBash {
		t.Skip("bash not found")
	}

	assert.NotEmpty(t, bash.Which())
	o, err := bash.Script("echo 'hello world'").Output()
	assert.NoError(t, err)
	assert.Equal(t, "hello world", strings.TrimSpace(o.Text()))
}
