package powershell

import (
	"fmt"
	"strings"

	"github.com/bearz-io/go/os/exec"
)

var (
	args = []string{"-NoLogo", "-NoProfile", "-NonInteractive", "-ExecutionPolicy", "Bypass"}
)

func init() {
	exec.Register("powershell", &exec.Executable{
		Name:     "powershell",
		Variable: "powershell_EXE",
		Windows: []string{
			"${WinDir}\\System32\\WindowsPowerShell\\v1.0\\powershell.exe",
			"${ProgramFiles}\\PowerShell\\7\\pwsh.exe",
			"${ProgramFiles(x86)}\\PowerShell\\7\\pwsh.exe",
			"${ProgramFiles}\\PowerShell\\6\\pwsh.exe",
			"${ProgramFiles(x86)}\\PowerShell\\6\\pwsh.exe",
		},
		Linux: []string{
			"/bin/pwsh",
			"/usr/bin/pwsh",
			"/usr/local/bin/pwsh",
			"/opt/microsoft/powershell/7/pwsh",
			"/opt/microsoft/powershell/6/pwsh",
		},
	})
}

func Which() *string {
	exe, _ := exec.Find("powershell", nil)
	if exe != "" {
		return &exe
	}

	return nil
}

func New(args ...string) *exec.Cmd {
	return exec.New("powershell", args...)
}

func File(file string) *exec.Cmd {
	splat := args[:]
	splat = append(splat, "-File", file)
	return New(splat...)
}

func Script(script string) *exec.Cmd {
	if !strings.ContainsAny(script, "\r\n;&|") {
		script = strings.TrimSpace(script)
		if strings.HasSuffix(script, ".ps1") {
			return File(script)
		}
	}

	tpl := `
$ErrorActionPreference = 'Stop'
%s 

if ($null -ne $Global:LASTEXITCODE) {
	exit $LASTEXITCODE
}
`
	script = fmt.Sprintf(tpl, script)

	splat := args[:]
	splat = append(splat, "-Command", script)
	return New(splat...)
}

func SetArgs(a ...string) {
	args = a
}
