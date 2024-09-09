package hostfile

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/bearz-io/go/errors"
	"github.com/bearz-io/go/fs"
	"github.com/bearz-io/go/os/env"
	"github.com/bearz-io/go/os/ps"
	"github.com/bearz-io/go/strings"
)

const TEST = "TEST"

func init() {
}

func GetPath() string {
	if runtime.GOOS == "windows" {
		winDir := env.Get("windir")
		if winDir == "" {
			winDir = "C:\\Windows"
		}

		return filepath.Join(winDir, "system32", "drivers", "etc", "hosts")
	}

	return "/etc/hosts"
}

func All() (map[string]string, error) {
	f, err := os.OpenFile(GetPath(), os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	sc := bufio.NewScanner(f)
	kv := make(map[string]string)
	for sc.Scan() {
		line := sc.Text()
		if strings.IsEmptySpace(line) {
			continue
		}

		if strings.HasPrefix(line, "#") {
			continue
		}

		line = strings.TrimSpace(line)
		parts := strings.Split(line, " ")
		if len(parts) == 2 {
			// not an ip
			if net.ParseIP(parts[0]) == nil {
				continue
			}

			kv[parts[1]] = parts[0]
		}
	}

	return kv, nil
}

func Has(cname string) (bool, error) {
	kv, err := All()
	if err != nil {
		return false, err
	}

	_, ok := kv[cname]
	return ok, nil
}

func HasIp(ip string) (bool, error) {
	kv, err := All()
	if err != nil {
		return false, err
	}

	for _, v := range kv {
		if v == ip {
			return true, nil
		}
	}

	return false, nil
}

func Remove(cname string) (bool, error) {
	if !ps.IsElevated() {
		return false, errors.ErrAccessDenied.WithResource(GetPath()).WithMessageF("access to hosts file requires elevated privileges")
	}

	err := Backup()
	if err != nil {
		return false, err
	}

	f, err := os.OpenFile(GetPath(), os.O_RDWR, os.ModePerm)
	if err != nil {
		return false, err
	}

	defer f.Close()

	sc := bufio.NewScanner(f)
	lines := []string{}
	updated := false

	for sc.Scan() {
		line := sc.Text()
		if strings.IsEmptySpace(line) {
			lines = append(lines, line)
			continue
		}

		if strings.HasPrefix(line, "#") {
			lines = append(lines, line)
			continue
		}

		line = strings.TrimSpace(line)
		parts := strings.Split(line, " ")
		if len(parts) == 2 {
			if parts[1] == cname {
				updated = true
				continue
			}
		}

		lines = append(lines, line)
	}

	if !updated {
		return false, nil
	}

	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		return false, err
	}

	size := 0
	for _, line := range lines {
		n, err := fmt.Fprintln(f, line)
		if err != nil {
			return false, err
		}

		size += n
	}

	err = f.Truncate(int64(size))
	if err != nil {
		return false, err
	}

	return true, nil
}

func Set(cname, ip string) error {

	if !ps.IsElevated() {
		return errors.ErrAccessDenied.WithResource(GetPath()).WithMessageF("access to hosts file requires elevated privileges")
	}

	err := Backup()
	if err != nil {
		return err
	}

	f, err := os.OpenFile(GetPath(), os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}

	defer f.Close()

	sc := bufio.NewScanner(f)
	lines := []string{}
	updated := false

	for sc.Scan() {
		line := sc.Text()
		if strings.IsEmptySpace(line) {
			lines = append(lines, line)
			continue
		}

		if strings.HasPrefix(line, "#") {
			lines = append(lines, line)
			continue
		}

		line = strings.TrimSpace(line)
		parts := strings.Split(line, " ")
		if len(parts) == 2 {
			if parts[1] == cname {
				updated = true
				continue
			}
		}

		ipParts := strings.Split(ip, ".")
		if len(ipParts) != 4 {
			return errors.Newf("invalid ip address: %s", ip)
		}

		lines = append(lines, fmt.Sprintf("%s %s", ip, cname))
	}

	if !updated {
		line := fmt.Sprintf("%s %s", strings.PadRight(ip, 20, " "), cname)
		lines = append(lines, line)
	}

	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}

	size := 0
	for _, line := range lines {
		n, err := fmt.Fprintln(f, line)
		if err != nil {
			return err
		}

		size += n
	}

	err = f.Truncate(int64(size))
	if err != nil {
		return err
	}

	return nil
}

func BackupAs(dest string) error {

	bytes, err := os.ReadFile(GetPath())
	if err != nil {
		return err
	}

	err = os.WriteFile(dest, bytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

func Backup() error {
	destDir := GetBackupDir()
	fs.EnsureDirDefault(destDir)
	dest := filepath.Join(destDir, fmt.Sprintf("hosts-%s.bak", time.Now().Format("2006-01-02-15-04-05")))
	return BackupAs(dest)
}

func GetBackupDir() string {
	return filepath.Join(os.TempDir(), "hosts-backups")
}

func RestoreFrom(src string) error {
	if !ps.IsElevated() {
		return errors.ErrAccessDenied.WithResource(GetPath()).WithMessageF("access to hosts file requires elevated privileges")
	}

	bytes, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	return os.WriteFile(GetPath(), bytes, 0644)
}
