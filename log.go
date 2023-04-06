package logd

import (
	"fmt"
	"io"
	"log"
	"os"
)

const (
	LINFO  = "INFO"
	LWARN  = "WARNING"
	LERR   = "ERROR"
	LFATAL = "FATAL"
)

type Logger struct {
	l    *log.Logger
	out  io.Writer
	fail func(msg ...interface{})
}

func NewLogger(out io.Writer, f func(...interface{}), flag int) *Logger {
	return &Logger{
		l:    log.New(out, "", flag),
		out:  out,
		fail: f,
	}
}

func (l *Logger) Close() error {
	out, ok := l.out.(io.Closer)
	if ok {
		return out.Close()
	}
	return nil
}

func (l *Logger) InfoPrintln(msg ...interface{}) {
	l.l.SetPrefix(fmt.Sprintf("%s: ", LINFO))
	err := l.l.Output(2, fmt.Sprintln(msg...))
	if err != nil {
		l.fail(err)
	}
}

func (l *Logger) WarnPrintln(msg ...interface{}) {
	l.l.SetPrefix(fmt.Sprintf("%s: ", LWARN))
	err := l.l.Output(2, fmt.Sprintln(msg...))
	if err != nil {
		l.fail(err)
	}
}

func (l *Logger) ErrPrintln(msg ...interface{}) {
	l.l.SetPrefix(fmt.Sprintf("%s: ", LERR))
	err := l.l.Output(2, fmt.Sprintln(msg...))
	if err != nil {
		l.fail(err)
	}
}

func (l *Logger) Fatal(msg ...interface{}) {
	l.l.SetPrefix(fmt.Sprintf("%s: ", LFATAL))
	l.l.Output(2, fmt.Sprintln(msg...))
	os.Exit(1)
}
