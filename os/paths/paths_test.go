package paths_test

import (
	"testing"

	"github.com/bearz-io/go/os/paths"
	"github.com/stretchr/testify/assert"
)

func TestPaths(t *testing.T) {
	home, err := paths.HomeDir()
	if err != nil {
		t.Errorf("Expected %v, got %v", nil, err)
	}

	assert.DirExists(t, home)
}
