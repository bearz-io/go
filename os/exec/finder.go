package exec

import (
	"runtime"

	"github.com/bearz-io/go/errors"
	"github.com/bearz-io/go/os/env"
	"github.com/bearz-io/go/runes"
	"github.com/bearz-io/go/strings"
)

type Executable struct {
	Name     string
	Path     string
	Variable string
	Windows  []string
	Unix     []string
	Darwin   []string
}

type ExecutableRegistry struct {
	data map[string]Executable
}

var Registry = &ExecutableRegistry{data: make(map[string]Executable)}

func (r *ExecutableRegistry) Register(name string, exe *Executable) {
	r.data[name] = *exe

	if exe.Variable == "" {
		sb := runes.Underscore([]rune(name), &runes.UnderscoreOptions{Screaming: true})
		exe.Variable = string(sb)
	}
}

func (r *ExecutableRegistry) Set(name string, exe *Executable) {
	r.data[name] = *exe
}

func (r *ExecutableRegistry) Get(name string) (*Executable, bool) {
	item, ok := r.data[name]
	return &item, ok
}

func (r *ExecutableRegistry) Has(name string) bool {
	_, ok := r.data[name]
	return ok
}

func (r *ExecutableRegistry) Find(name string, options *WhichOptions) (string, error) {
	m, ok := r.data[name]
	if !ok {
		sb := runes.Underscore([]rune(name), &runes.UnderscoreOptions{Screaming: true})
		m = Executable{Name: name}
		m.Variable = string(sb)
		r.data[name] = m
	}

	if options == nil {
		options = &WhichOptions{}
	}

	if options.UseCache && m.Path != "" {
		return m.Path, nil
	}

	if m.Variable != "" {
		value := env.Get(m.Variable)
		if value != "" {
			value = env.ExpandSafe(value)
			if value != "" {
				next, ok := WhichFirst(value, options)
				if ok {
					m.Path = next
					return m.Path, nil
				}
			}
		}
	}

	if m.Path != "" {
		next, ok := WhichFirst(m.Path, options)
		if ok {
			m.Path = next
			return m.Path, nil
		}
	}

	if runtime.GOOS == "windows" {
		for _, path := range m.Windows {
			if strings.IsEmptySpace(path) {
				continue
			}

			exe2 := env.ExpandSafe(path)
			if exe2 == "" {
				continue
			}

			next, ok := WhichFirst(exe2, options)
			if ok {
				m.Path = next
				return m.Path, nil
			}
		}

		return "", errors.ErrNotFound.WithMessageF("executable not found: %s", name)
	}

	if runtime.GOOS == "darwin" {
		for _, path := range m.Darwin {
			if strings.IsEmptySpace(path) {
				continue
			}

			exe2 := env.ExpandSafe(path)
			if exe2 == "" {
				continue
			}

			next, ok := WhichFirst(exe2, options)
			if ok {
				m.Path = next
				return m.Path, nil
			}
		}

		// fallthrough to unix
	}

	for _, path := range m.Darwin {
		if strings.IsEmptySpace(path) {
			continue
		}

		exe2 := env.ExpandSafe(path)
		if exe2 == "" {
			continue
		}

		next, ok := WhichFirst(exe2, options)
		if ok {
			m.Path = next
			return m.Path, nil
		}
	}

	return "", errors.ErrNotFound.WithMessageF("executable not found: %s", name)
}

func Register(name string, exe *Executable) {
	Registry.Register(name, exe)
}

func Find(name string, options *WhichOptions) (string, error) {
	return Registry.Find(name, options)
}
