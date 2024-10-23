package fs

import (
	"io/fs"
	"os"
	"path/filepath"
)

func init() {
}

func Create(path string) (*os.File, error) {
	return os.Create(path)
}

func Exists(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}

func Mkdir(path string, perm fs.FileMode) error {
	return os.Mkdir(path, perm)
}

func EnsureDir(path string, perm fs.FileMode) error {
	if !Exists(path) {
		err := os.MkdirAll(path, perm)
		if err != nil {
			return err
		}
	}

	return nil
}

func EnsureDirDefault(path string) error {
	return EnsureDir(path, 0755)
}

func EnsureFile(path string, perm fs.FileMode) error {
	if !Exists(path) {
		fi, err := os.Create(path)
		if err != nil {
			return err
		}

		fi.Chmod(perm)
		return fi.Close()
	}

	return nil
}

func EnsureFileDefault(path string) error {
	return EnsureFile(path, 0644)
}

func Resolve(path string) (string, error) {
	if path == "" {
		return ".", nil
	}

	if path[0] == '~' {
		home := os.Getenv("HOME")
		if home == "" {
			return "", os.ErrNotExist
		}

		return filepath.Join(home, path[1:]), nil
	}

	if !filepath.IsAbs(path) {
		p1, err := filepath.Abs(path)
		if err != nil {
			return "", err
		}

		return p1, nil
	} else {
		path = filepath.Clean(path)
	}

	return path, nil
}

func OpenRead(path string, perm fs.FileMode) (*os.File, error) {
	p, err := Resolve(path)
	if err != nil {
		return nil, err
	}

	return os.OpenFile(p, os.O_RDONLY, perm)
}

func OpenReadDefault(path string) (*os.File, error) {
	return OpenRead(path, 0644)
}

func OpenWrite(path string, perm fs.FileMode) (*os.File, error) {
	p, err := Resolve(path)
	if err != nil {
		return nil, err
	}

	return os.OpenFile(p, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
}

func OpenWriteDefault(path string) (*os.File, error) {
	return OpenWrite(path, 0644)
}
