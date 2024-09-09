package bash

import (
	"path/filepath"
	"runtime"
	"strings"
	"unicode"

	"github.com/bearz-io/go/fs"
	"github.com/bearz-io/go/os/env"
	"github.com/bearz-io/go/os/exec"
)

const TEST = "TEST"

var (
	wslInstalled = false
	args         = []string{"-noprofile", "--norc", "-e", "-o", "pipefail"}
)

func init() {
	exec.Register("bash", &exec.Executable{
		Name:     "bash",
		Variable: "BASH_EXE",
		Windows: []string{
			"${ProgramFiles}\\Git\\bin\\bash.exe",
			"${ProgramFiles(x86)}\\Git\\bin\\bash.exe",
			"${ProgramFiles}\\Git\\usr\\bin\\bash.exe",
			"${ProgramFiles(x86)}\\Git\\usr\\bin\\bash.exe",
			"${ChocolateyInstall}\\msys2\\usr\\bin\\bash.exe",
			"${windir}\\System32\\bash.exe",
			"${LocalAppData}\\Microsoft\\WindowsApps\\bash.exe",
		},
		Linux: []string{
			"/bin/bash",
			"/usr/bin/bash",
			"/usr/local/bin/bash",
		},
	})

	if runtime.GOOS == "windows" {
		path2, err := env.Expand("${LocalAppData}\\Microsoft\\WindowsApps\\bash.exe", nil)
		if err == nil {
			wslInstalled = fs.Exists(path2)
		} else {
			path3, err := env.Expand("${windir}\\System32\\bash.exe", nil)
			if err == nil {
				wslInstalled = fs.Exists(path3)
			}
		}
	}
}

// finds which bash executabile will be used by
// `exec.Find` and `exec.Command`
func Which() *string {
	exe, _ := exec.Find("bash", nil)
	if exe != "" {
		return &exe
	}

	return nil
}

// Creates a new command where the path is the bash executable
// and the args are set.
func New(args ...string) *exec.Cmd {
	return exec.New("bash", args...)
}

func File(path string) *exec.Cmd {
	splat := args[:]
	if wslInstalled {
		exe, _ := exec.Find("bash", nil)
		if exe != "" {
			if strings.HasSuffix("System32\\base.exe", exe) || strings.HasSuffix("WindowsApps\\bash.exe", exe) {
				if !filepath.IsAbs(path) {
					f, err := filepath.Abs(path)
					if err == nil {
						path = f
					}

					path = "/mnt/" + string(unicode.ToLower(rune(path[0]))) + path[2:]
					path = filepath.ToSlash(path)
				}
			}
		}
	}

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

// Changes the default args for File and Script.  The
// default args are ["-noprofile", "--norc", "-e", "-o", "pipefail"]
func SetArgs(a ...string) {
	args = a
}
