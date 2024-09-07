package fs

import "os"

func init() {
}

func Exists(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}
