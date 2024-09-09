package powershell_test

import (
	"strings"
	"testing"

	"github.com/bearz-io/go/spawn/powershell"
	"github.com/stretchr/testify/assert"
)

func TestPowershell(t *testing.T) {
	hasPwsh := powershell.Which() != nil
	if !hasPwsh {
		t.Skip("pwsh not found")
	}

	assert.NotEmpty(t, powershell.Which())
	o, err := powershell.Script("echo 'hello world'").Output()
	assert.NoError(t, err)
	assert.Equal(t, "hello world", strings.TrimSpace(o.Text()))
}
