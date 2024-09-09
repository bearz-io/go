package pwsh_test

import (
	"strings"
	"testing"

	"github.com/bearz-io/go/spawn/pwsh"
	"github.com/stretchr/testify/assert"
)

func TestPwsh(t *testing.T) {
	hasPwsh := pwsh.Which() != nil
	if !hasPwsh {
		t.Skip("pwsh not found")
	}

	assert.NotEmpty(t, pwsh.Which())
	o, err := pwsh.Script("echo 'hello world'").Output()
	assert.NoError(t, err)
	assert.Equal(t, "hello world", strings.TrimSpace(o.Text()))
}
