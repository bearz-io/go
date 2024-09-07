//go:build aix || darwin || dragonfly || freebsd || hurd || illumos || ios || linux || netbsd || openbsd || plan9 || solaris || zos

package os

const (
	POSIX    = true
	EOL      = "\n"
	PATH_SEP = ":"
	DIR_SEP  = "/"
	DEV_NULL = "/dev/null"
)
