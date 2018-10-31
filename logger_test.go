package golorich_test

import (
	"bytes"
	"io"
	"log"
	"os"
	"os/exec"
	"reflect"
	"regexp"
	"testing"

	"github.com/dip-dev/golorich"
	"github.com/hashicorp/logutils"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "success",
			want: "*golorich.Logger",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := golorich.New(&bytes.Buffer{}, "", log.LstdFlags, golorich.Debug)
			ts := reflect.TypeOf(got).String()
			if tt.want != ts {
				t.Errorf("want %s, but %s", tt.want, ts)
			}
		})
	}
}

func TestGetLevelFromString(t *testing.T) {
	tests := []struct {
		name  string
		level string
		want  logutils.LogLevel
	}{
		{
			name:  "all_lower_case",
			level: "debug",
			want:  golorich.Debug,
		},
		{
			name:  "pascal_case",
			level: "Info",
			want:  golorich.Info,
		},
		{
			name:  "all_upper_case",
			level: "WARN",
			want:  golorich.Warn,
		},
		{
			name:  "nonexistent_level",
			level: "invalid",
			want:  golorich.Info,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := golorich.GetLevelFromString(tt.level)
			if tt.want != got {
				t.Fatalf("want %s, but %s", tt.want, got)
			}
		})
	}
}

func TestLogger_Debugf(t *testing.T) {
	tests := []struct {
		name   string
		buf    *bytes.Buffer
		format string
		args   []interface{}
		want   string
	}{
		{
			name:   "success",
			buf:    &bytes.Buffer{},
			format: "%s %s",
			args:   []interface{}{"test", "message"},
			want:   `^\d{4}\/\d{2}\/\d{2} \d{2}:\d{2}:\d{2} \[DEBUG\] test message\n$`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger := golorich.New(tt.buf, "", log.LstdFlags, golorich.Debug)
			logger.Debugf(tt.format, tt.args...)
			got := tt.buf.String()
			if ok, err := regexp.MatchString(tt.want, got); err != nil {
				t.Fatalf("%v", err)
			} else if !ok {
				t.Errorf("not match pattern: want %s, but %s", tt.want, got)
			}
		})
	}
}

func TestLogger_Debugln(t *testing.T) {
	tests := []struct {
		name string
		buf  *bytes.Buffer
		args []interface{}
		want string
	}{
		{
			name: "success",
			buf:  &bytes.Buffer{},
			args: []interface{}{"test", "message"},
			want: `^\d{4}\/\d{2}\/\d{2} \d{2}:\d{2}:\d{2} \[DEBUG\] test message\n$`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger := golorich.New(tt.buf, "", log.LstdFlags, golorich.Debug)
			logger.Debugln(tt.args...)
			got := tt.buf.String()
			if ok, err := regexp.MatchString(tt.want, got); err != nil {
				t.Fatalf("%v", err)
			} else if !ok {
				t.Errorf("not match pattern: want %s, but %s", tt.want, got)
			}
		})
	}
}

func TestLogger_Infof(t *testing.T) {
	tests := []struct {
		name   string
		buf    *bytes.Buffer
		format string
		args   []interface{}
		want   string
	}{
		{
			name:   "success",
			buf:    &bytes.Buffer{},
			format: "%s %s",
			args:   []interface{}{"test", "message"},
			want:   `^\d{4}\/\d{2}\/\d{2} \d{2}:\d{2}:\d{2} \[INFO\] test message\n$`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger := golorich.New(tt.buf, "", log.LstdFlags, golorich.Info)
			logger.Infof(tt.format, tt.args...)
			got := tt.buf.String()
			if ok, err := regexp.MatchString(tt.want, got); err != nil {
				t.Fatalf("%v", err)
			} else if !ok {
				t.Errorf("not match pattern: want %s, but %s", tt.want, got)
			}
		})
	}
}

func TestLogger_Infoln(t *testing.T) {
	tests := []struct {
		name string
		buf  *bytes.Buffer
		args []interface{}
		want string
	}{
		{
			name: "success",
			buf:  &bytes.Buffer{},
			args: []interface{}{"test", "message"},
			want: `^\d{4}\/\d{2}\/\d{2} \d{2}:\d{2}:\d{2} \[INFO\] test message\n$`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger := golorich.New(tt.buf, "", log.LstdFlags, golorich.Info)
			logger.Infoln(tt.args...)
			got := tt.buf.String()
			if ok, err := regexp.MatchString(tt.want, got); err != nil {
				t.Fatalf("%v", err)
			} else if !ok {
				t.Errorf("not match pattern: want %s, but %s", tt.want, got)
			}
		})
	}
}

func TestLogger_Warnf(t *testing.T) {
	tests := []struct {
		name   string
		buf    *bytes.Buffer
		format string
		args   []interface{}
		want   string
	}{
		{
			name:   "success",
			buf:    &bytes.Buffer{},
			format: "%s %s",
			args:   []interface{}{"test", "message"},
			want:   `^\d{4}\/\d{2}\/\d{2} \d{2}:\d{2}:\d{2} \[WARN\] test message\n$`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger := golorich.New(tt.buf, "", log.LstdFlags, golorich.Warn)
			logger.Warnf(tt.format, tt.args...)
			got := tt.buf.String()
			if ok, err := regexp.MatchString(tt.want, got); err != nil {
				t.Fatalf("%v", err)
			} else if !ok {
				t.Errorf("not match pattern: want %s, but %s", tt.want, got)
			}
		})
	}
}

func TestLogger_Warnln(t *testing.T) {
	tests := []struct {
		name string
		buf  *bytes.Buffer
		args []interface{}
		want string
	}{
		{
			name: "success",
			buf:  &bytes.Buffer{},
			args: []interface{}{"test", "message"},
			want: `^\d{4}\/\d{2}\/\d{2} \d{2}:\d{2}:\d{2} \[WARN\] test message\n$`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger := golorich.New(tt.buf, "", log.LstdFlags, golorich.Warn)
			logger.Warnln(tt.args...)
			got := tt.buf.String()
			if ok, err := regexp.MatchString(tt.want, got); err != nil {
				t.Fatalf("%v", err)
			} else if !ok {
				t.Errorf("not match pattern: want %s, but %s", tt.want, got)
			}
		})
	}
}

func TestLogger_Errorf(t *testing.T) {
	tests := []struct {
		name   string
		buf    *bytes.Buffer
		format string
		args   []interface{}
		want   string
	}{
		{
			name:   "success",
			buf:    &bytes.Buffer{},
			format: "%s %s",
			args:   []interface{}{"test", "message"},
			want:   `^\d{4}\/\d{2}\/\d{2} \d{2}:\d{2}:\d{2} \[ERROR\] test message\n$`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger := golorich.New(tt.buf, "", log.LstdFlags, golorich.Error)
			logger.Errorf(tt.format, tt.args...)
			got := tt.buf.String()
			if ok, err := regexp.MatchString(tt.want, got); err != nil {
				t.Fatalf("%v", err)
			} else if !ok {
				t.Errorf("not match pattern: want %s, but %s", tt.want, got)
			}
		})
	}
}

func TestLogger_Errorln(t *testing.T) {
	tests := []struct {
		name string
		buf  *bytes.Buffer
		args []interface{}
		want string
	}{
		{
			name: "success",
			buf:  &bytes.Buffer{},
			args: []interface{}{"test", "message"},
			want: `^\d{4}\/\d{2}\/\d{2} \d{2}:\d{2}:\d{2} \[ERROR\] test message\n$`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger := golorich.New(tt.buf, "", log.LstdFlags, golorich.Error)
			logger.Errorln(tt.args...)
			got := tt.buf.String()
			if ok, err := regexp.MatchString(tt.want, got); err != nil {
				t.Fatalf("%v", err)
			} else if !ok {
				t.Errorf("not match pattern: want %s, but %s", tt.want, got)
			}
		})
	}
}

func TestLogger_Fatalf(t *testing.T) {
	tests := []struct {
		name   string
		out    io.Writer
		format string
		args   []interface{}
		want   string
	}{
		{
			name:   "success",
			out:    os.Stdout,
			format: "%s %s",
			args:   []interface{}{"test", "message"},
			want:   `^\d{4}\/\d{2}\/\d{2} \d{2}:\d{2}:\d{2} \[FATAL\] test message\n$`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger := golorich.New(tt.out, "", log.LstdFlags, golorich.Fatal)
			if os.Getenv("BE_FATAL") == "1" {
				logger.Fatalf(tt.format, tt.args...)
				return
			}
			cmd := exec.Command(os.Args[0], "-test.run=TestLogger_Fatalf")
			cmd.Env = append(os.Environ(), "BE_FATAL=1")
			got, err := cmd.Output()
			if ok, err := regexp.MatchString(tt.want, string(got)); err != nil {
				t.Fatalf("%v", err)
			} else if !ok {
				t.Errorf("not match pattern: want %s, but %s", tt.want, got)
			}
			if e, ok := err.(*exec.ExitError); ok && !e.Success() {
				return
			}
			t.Fatalf("process ran with err %v, want os.Exit(1)", err)
		})
	}
}

func TestLogger_Fatalln(t *testing.T) {
	tests := []struct {
		name string
		out  io.Writer
		args []interface{}
		want string
	}{
		{
			name: "success",
			out:  os.Stdout,
			args: []interface{}{"test", "message"},
			want: `^\d{4}\/\d{2}\/\d{2} \d{2}:\d{2}:\d{2} \[FATAL\] test message\n$`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger := golorich.New(tt.out, "", log.LstdFlags, golorich.Fatal)
			if os.Getenv("BE_FATAL") == "1" {
				logger.Fatalln(tt.args...)
				return
			}
			cmd := exec.Command(os.Args[0], "-test.run=TestLogger_Fatalln")
			cmd.Env = append(os.Environ(), "BE_FATAL=1")
			got, err := cmd.Output()
			if ok, err := regexp.MatchString(tt.want, string(got)); err != nil {
				t.Fatalf("%v", err)
			} else if !ok {
				t.Errorf("not match pattern: want %s, but %s", tt.want, got)
			}
			if e, ok := err.(*exec.ExitError); ok && !e.Success() {
				return
			}
			t.Fatalf("process ran with err %v, want os.Exit(1)", err)
		})
	}
}
