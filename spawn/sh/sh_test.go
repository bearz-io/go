package sh_test

import (
	"strings"
	"testing"

	"github.com/bearz-io/go/spawn/sh"
	"github.com/stretchr/testify/assert"
)

func TestSh(t *testing.T) {
	hasSh := sh.Which() != nil
	if !hasSh {
		t.Skip("sh not found")
	}

	assert.NotEmpty(t, sh.Which())
	o, err := sh.Script("echo 'hello world'").Output()
	assert.NoError(t, err)
	assert.Equal(t, "hello world", strings.TrimSpace(o.Text()))
}
