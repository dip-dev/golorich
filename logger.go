package golorich

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/hashicorp/logutils"
)

// ...
const (
	Debug logutils.LogLevel = "DEBUG"
	Info  logutils.LogLevel = "INFO"
	Warn  logutils.LogLevel = "WARN"
	Error logutils.LogLevel = "ERROR"
	Fatal logutils.LogLevel = "FATAL"
)

var logLevels = []logutils.LogLevel{
	Debug,
	Info,
	Warn,
	Error,
	Fatal,
}

const (
	debugPrefix = "[DEBUG] "
	infoPrefix  = "[INFO] "
	warnPrefix  = "[WARN] "
	errorPrefix = "[ERROR] "
	fatalPrefix = "[FATAL] "
)

// Logger ...
type Logger struct {
	*log.Logger
}

// New ...
func New(out io.Writer, prefix string, flag int, minLevel logutils.LogLevel) *Logger {
	filter := &logutils.LevelFilter{
		Levels:   logLevels,
		MinLevel: minLevel,
		Writer:   out,
	}
	logger := log.New(filter, prefix, flag)
	return &Logger{Logger: logger}
}

// GetLevelFromString gets the log level from string like 'debug', 'Info', 'WARN', and so on.
func GetLevelFromString(minLevel string) logutils.LogLevel {
	var ret = Info
	for _, level := range logLevels {
		lowerLevel := string(level)
		if strings.ToUpper(minLevel) == lowerLevel {
			ret = logutils.LogLevel(level)
		}
	}
	return ret
}

// Debugf logs to the DEBUG log in fmt.Printf manner.
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.Output(2, fmt.Sprint(debugPrefix, fmt.Sprintf(format, v...)))
}

// Debugln logs to the DEBUG log in fmt.Println manner.
func (l *Logger) Debugln(v ...interface{}) {
	l.Output(2, debugPrefix+fmt.Sprintln(v...))
}

// Infof logs to the INFO log in fmt.Printf manner.
func (l *Logger) Infof(format string, v ...interface{}) {
	l.Output(2, fmt.Sprint(infoPrefix, fmt.Sprintf(format, v...)))
}

// Infoln logs to the INFO log in fmt.Println manner.
func (l *Logger) Infoln(v ...interface{}) {
	l.Output(2, infoPrefix+fmt.Sprintln(v...))
}

// Warnf logs to the WARN log in fmt.Printf manner.
func (l *Logger) Warnf(format string, v ...interface{}) {
	l.Output(2, fmt.Sprint(warnPrefix, fmt.Sprintf(format, v...)))
}

// Warnln logs to the WARN log in fmt.Println manner.
func (l *Logger) Warnln(v ...interface{}) {
	l.Output(2, warnPrefix+fmt.Sprintln(v...))
}

// Errorf logs to the ERROR log in fmt.Printf manner.
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.Output(2, fmt.Sprint(errorPrefix, fmt.Sprintf(format, v...)))
}

// Errorln logs to the ERROR log in fmt.Println manner.
func (l *Logger) Errorln(v ...interface{}) {
	l.Output(2, errorPrefix+fmt.Sprintln(v...))
}

// Fatalf logs to the FATAL log in fmt.Printf manner.
func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.Output(2, fmt.Sprint(fatalPrefix, fmt.Sprintf(format, v...)))
	os.Exit(1)
}

// Fatalln logs to the FATAL log in fmt.Println manner.
func (l *Logger) Fatalln(v ...interface{}) {
	l.Output(2, fatalPrefix+fmt.Sprintln(v...))
	os.Exit(1)
}
