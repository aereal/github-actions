package toolkit

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func SetOutput(name, value string) {
	fmt.Fprintf(out, "::set-output name=%s::%s\n", name, value)
}

func SetEnv(name, value string) {
	fmt.Fprintf(out, "::set-env name=%s::%s\n", name, value)
}

func AddPath(paths ...string) {
	for _, path := range paths {
		fmt.Fprintf(out, "::add-path::%s\n", path)
	}
}

type WarningOptions struct {
	File   string
	Line   int
	Column int
}

func Warning(msg string, opts *WarningOptions) {
	fmt.Fprint(out, "::warning")
	if opts != nil {
		fmt.Fprintf(out, " file=%s,line=%d,col=%d", opts.File, opts.Line, opts.Column)
	}
	fmt.Fprintf(out, "::%s\n", msg)
}

type ErrorOptions struct {
	File   string
	Line   int
	Column int
}

func Error(msg string, opts *ErrorOptions) {
	fmt.Fprint(out, "::error")
	if opts != nil {
		fmt.Fprintf(out, " file=%s,line=%d,col=%d", opts.File, opts.Line, opts.Column)
	}
	fmt.Fprintf(out, "::%s\n", msg)
}
