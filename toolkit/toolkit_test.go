package toolkit

import (
	"bytes"
	"testing"
)

func init() {
	out = new(bytes.Buffer)
}

func TestSetOutput(t *testing.T) {
	type args struct {
		name  string
		value string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ok",
			args: args{
				name:  "poppoe",
				value: "ubobo",
			},
			want: "::set-output name=poppoe::ubobo\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reset()
			SetOutput(tt.args.name, tt.args.value)
			if got := getLastOutput(); got != tt.want {
				t.Errorf("SetOutput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetEnv(t *testing.T) {
	type args struct {
		name  string
		value string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ok",
			args: args{
				name:  "GO111MODULE",
				value: "1",
			},
			want: "::set-env name=GO111MODULE::1\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reset()
			SetEnv(tt.args.name, tt.args.value)
			if got := getLastOutput(); got != tt.want {
				t.Errorf("SetEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddPath(t *testing.T) {
	type args struct {
		paths []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "add a path",
			args: args{
				paths: []string{"/home/runner/bin"},
			},
			want: "::add-path::/home/runner/bin\n",
		},
		{
			name: "add multiple paths",
			args: args{
				paths: []string{"/home/runner/bin", "/home/runner/sbin"},
			},
			want: "::add-path::/home/runner/bin\n::add-path::/home/runner/sbin\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reset()
			AddPath(tt.args.paths...)
			if got := getLastOutput(); got != tt.want {
				t.Errorf("AddPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWarning(t *testing.T) {
	type args struct {
		msg  string
		opts *WarningOptions
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "only message",
			args: args{
				msg: "poyoyo",
			},
			want: "::warning::poyoyo\n",
		},
		{
			name: "with options",
			args: args{
				msg: "poyoyo",
				opts: &WarningOptions{
					File:   "main.go",
					Line:   3,
					Column: 4,
				},
			},
			want: "::warning file=main.go,line=3,col=4::poyoyo\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reset()
			Warning(tt.args.msg, tt.args.opts)
			if got := getLastOutput(); got != tt.want {
				t.Errorf("Warning() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestError(t *testing.T) {
	type args struct {
		msg  string
		opts *ErrorOptions
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "only message",
			args: args{
				msg: "poyoyo",
			},
			want: "::error::poyoyo\n",
		},
		{
			name: "with options",
			args: args{
				msg: "poyoyo",
				opts: &ErrorOptions{
					File:   "main.go",
					Line:   3,
					Column: 4,
				},
			},
			want: "::error file=main.go,line=3,col=4::poyoyo\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reset()
			Error(tt.args.msg, tt.args.opts)
			if got := getLastOutput(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func reset() {
	out.(*bytes.Buffer).Reset()
}

func getLastOutput() string {
	return out.(*bytes.Buffer).String()
}
