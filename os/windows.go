//go:build windows

package os

const (
	FAMILY   = "windows"
	POSIX    = false
	EOL      = "\r\n"
	PATH_SEP = ";"
	DIR_SEP  = "\\"
	DEV_NULL = "NUL"
)
