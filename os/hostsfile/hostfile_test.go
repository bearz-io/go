package hostfile_test

import (
	"testing"

	hostfile "github.com/bearz-io/go/os/hostsfile"
	"github.com/stretchr/testify/assert"
)

func TestHostfile(t *testing.T) {
	assert.Equal(t, hostfile.TEST, "TEST")
}
