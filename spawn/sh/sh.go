package sh

import (
	"strings"

	"github.com/bearz-io/go/os/exec"
)

var (
	args = []string{"-e"}
)

func init() {
	exec.Register("sh", &exec.Executable{
		Name:     "sh",
		Variable: "SH_EXE",
		Windows: []string{
			"${ProgramFiles}\\Git\\usr\\bin\\sh.exe",
			"${ProgramFiles(x86)}\\Git\\usr\\bin\\sh.exe",
			"${ChocolateyInstall}\\msys2\\usr\\bin\\sh.exe",
		},
		Linux: []string{
			"/bin/sh",
			"/usr/bin/sh",
			"/usr/local/bin/sh",
		},
	})
}

func Which() *string {
	exe, _ := exec.Find("sh", nil)
	if exe != "" {
		return &exe
	}

	return nil
}

func New(args ...string) *exec.Cmd {
	return exec.New("sh", args...)
}

func File(path string) *exec.Cmd {
	splat := args[:]
	splat = append(splat, path)
	return New(splat...)
}

func Script(script string) *exec.Cmd {
	if !strings.ContainsAny(script, "\r\n;&|") {
		script = strings.TrimSpace(script)
		if strings.HasSuffix(script, ".sh") {
			return File(script)
		}
	}

	splat := args[:]
	splat = append(splat, "-c", script)
	return New(splat...)
}

func SetArgs(a ...string) {
	args = a
}
