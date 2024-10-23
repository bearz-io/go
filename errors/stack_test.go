package errors_test

import (
	"fmt"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"testing"

	"github.com/bearz-io/errors"
)

const (
	prefix     = "  github.com/bearz-io/errors"
	testPrefix = "  github.com/bearz-io/errors_test"
)

var (
	initpc = caller()
	cwd    string
)

func init() {
	x := X{}
	f := x.ptr().File()
	cwd = filepath.Dir(f)
}

type X struct{}

// val returns a Frame pointing to itself.
func (x X) val() errors.Frame {
	return caller()
}

// ptr returns a Frame pointing to itself.
func (x *X) ptr() errors.Frame {
	return caller()
}

func TestFrameFormat(t *testing.T) {
	var tests = []struct {
		errors.Frame
		format string
		want   string
	}{{
		initpc,
		"%s",
		"stack_test.go",
	}, {
		initpc,
		"%+s",
		testPrefix + ".init\n" +
			"    at " + cwd + "/stack_test.go",
	}, {
		0,
		"%s",
		"unknown",
	}, {
		0,
		"%+s",
		"unknown",
	}, {
		initpc,
		"%d",
		"20",
	}, {
		0,
		"%d",
		"0",
	}, {
		initpc,
		"%n",
		"init",
	}, {
		func() errors.Frame {
			var x X
			return x.ptr()
		}(),
		"%n",
		`\(\*X\).ptr`,
	}, {
		func() errors.Frame {
			var x X
			return x.val()
		}(),
		"%n",
		"X.val",
	}, {
		0,
		"%n",
		"",
	}, {
		initpc,
		"%v",
		"stack_test.go:20",
	}, {
		initpc,
		"%+v",
		testPrefix + ".init\n" +
			"    at " + cwd + "/stack_test.go:20",
	}, {
		0,
		"%v",
		"unknown:0",
	}}

	for i, tt := range tests {
		testFormatRegexp(t, i, tt.Frame, tt.format, tt.want)
	}
}

func TestFuncname(t *testing.T) {
	tests := []struct {
		name, want string
	}{
		{"", ""},
		{"runtime.main", "main"},
		{"github.com/pkg/errors.funcname", "funcname"},
		{"funcname", "funcname"},
		{"io.copyBuffer", "copyBuffer"},
		{"main.(*R).Write", "(*R).Write"},
	}

	for _, tt := range tests {
		got := funcname(tt.name)
		want := tt.want
		if got != want {
			t.Errorf("funcname(%q): want: %q, got %q", tt.name, want, got)
		}
	}
}

func TestStackTrace(t *testing.T) {
	tests := []struct {
		err  error
		want []string
	}{{
		errors.New("ooh"), []string{
			"  github.com/bearz-io/errors_test.TestStackTrace\n" +
				"    at " + cwd + "/stack_test.go:140",
		},
	}, {
		errors.Wrap(errors.New("ooh"), "ahh"), []string{
			"  github.com/bearz-io/errors_test.TestStackTrace\n" +
				"    at " + cwd + "/stack_test.go:145", // this is the stack of Wrap, not New
		},
	}, /*{
			errors.Cause(errors.Wrap(errors.New("ooh"), "ahh")), []string{
				"  github.com/bearz-io/go-errors_test.TestStackTrace\n" +
					"    at " + cwd + "/stack_test.go:131", // this is the stack of New
			},
		},*/{
			func() error { return errors.New("ooh") }(), []string{
				`  github.com/bearz-io/errors_test.TestStackTrace.func1` +
					"\n    at " + cwd + "/stack_test.go:155", // this is the stack of New
				"  github.com/bearz-io/errors_test.TestStackTrace\n" +
					"    at " + cwd + "/stack_test.go:155", // this is the stack of New's caller
			},
		}} /*, {
		errors.Cause(func() error {
			return func() error {
				return errors.Errorf("hello %s", fmt.Sprintf("world: %s", "ooh"))
			}()
		}()), []string{
			`   github.com/bearz-io/go-errors_test.TestStackTrace.func2.1` +
				"\n    at " + cwd + "/stack_test.go:145", // this is the stack of Errorf
			`github.com/bearz-io/go-errors_test.TestStackTrace.func2` +
				"\n    at " + cwd + "/stack_test.go:146", // this is the stack of Errorf's caller
			"github.com/bearz-io/go-errors_test.TestStackTrace\n" +
				"\t    at " + cwd + "/stack_test.go:147", // this is the stack of Errorf's caller's caller
		},
	}}*/
	for i, tt := range tests {
		x, ok := tt.err.(interface {
			StackTrace() errors.StackTrace
		})
		if !ok {
			t.Errorf("expected %#v to implement StackTrace() StackTrace", tt.err)
			continue
		}
		st := x.StackTrace()
		for j, want := range tt.want {
			testFormatRegexp(t, i, st[j], "%+v", want)
		}
	}
}

func TestStackTraceFormat(t *testing.T) {
	tests := []struct {
		errors.StackTrace
		format string
		want   string
	}{{
		nil,
		"%s",
		`\[\]`,
	}, {
		nil,
		"%v",
		`\[\]`,
	}, {
		nil,
		"%+v",
		"",
	}, {
		nil,
		"%#v",
		`\[\]errors.Frame\(nil\)`,
	}, {
		make(errors.StackTrace, 0),
		"%s",
		`\[\]`,
	}, {
		make(errors.StackTrace, 0),
		"%v",
		`\[\]`,
	}, {
		make(errors.StackTrace, 0),
		"%+v",
		"",
	}, {
		make(errors.StackTrace, 0),
		"%#v",
		`\[\]errors.Frame{}`,
	}}
	/*{
		stackTrace()[:2],
		"%s",
		`\[stack_test.go stack_test.go\]`,
	}, {
		stackTrace()[:2],
		"%v",
		`\[stack_test.go:174 stack_test.go:221\]`,
	}, {
		stackTrace()[:2],
		"%+v",
		"\n" +
			"github.com/pkg/errors.stackTrace\n" +
			"\t.+/github.com/pkg/errors/stack_test.go:174\n" +
			"github.com/pkg/errors.TestStackTraceFormat\n" +
			"\t.+/github.com/pkg/errors/stack_test.go:225",
	}, {
		stackTrace()[:2],
		"%#v",
		`\[\]errors.Frame{stack_test.go:174, stack_test.go:233}`,
	} }
	*/
	for i, tt := range tests {
		testFormatRegexp(t, i, tt.StackTrace, tt.format, tt.want)
	}
}

// a version of runtime.Caller that returns a Frame, not a uintptr.
func caller() errors.Frame {
	var pcs [3]uintptr
	n := runtime.Callers(2, pcs[:])
	frames := runtime.CallersFrames(pcs[:n])
	frame, _ := frames.Next()
	return errors.Frame(frame.PC)
}

// funcname removes the path prefix component of a function's name reported by func.Name().
func funcname(name string) string {
	i := strings.LastIndex(name, "/")
	name = name[i+1:]
	i = strings.Index(name, ".")
	return name[i+1:]
}

func testFormatRegexp(t *testing.T, n int, arg interface{}, format, want string) {
	t.Helper()
	got := fmt.Sprintf(format, arg)
	gotLines := strings.SplitN(got, "\n", -1)
	wantLines := strings.SplitN(want, "\n", -1)

	if len(wantLines) > len(gotLines) {
		t.Errorf("test %d: wantLines(%d) > gotLines(%d):\n got: %q\nwant: %q", n+1, len(wantLines), len(gotLines), got, want)
		return
	}

	for i, w := range wantLines {
		match, err := regexp.MatchString(w, gotLines[i])
		if err != nil {
			t.Fatal(err)
		}
		if !match {
			t.Errorf("test %d: line %d: fmt.Sprintf(%q, err):\n got: %q\nwant: %q", n+1, i+1, format, got, want)
		}
	}
}
