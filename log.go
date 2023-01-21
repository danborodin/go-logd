package main

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
	l   *log.Logger
	out io.Writer
}

func NewLogger(out io.Writer, flag int) *Logger {
	return &Logger{
		l:   log.New(out, "", flag),
		out: out,
	}
}

func (l *Logger) Close() error {
	file, ok := l.out.(*os.File)
	if ok {
		return file.Close()
	}
	return nil
}

func (l *Logger) InfoPrintln(msg ...interface{}) error {
	l.l.SetPrefix(fmt.Sprintf("%s ", LINFO))
	return l.l.Output(2, fmt.Sprintln(msg...))
}

func (l *Logger) WarnPrintln(msg ...interface{}) error {
	l.l.SetPrefix(fmt.Sprintf("%s ", LWARN))
	return l.l.Output(2, fmt.Sprintln(msg...))
}

func (l *Logger) ErrPrintln(msg ...interface{}) error {
	l.l.SetPrefix(fmt.Sprintf("%s ", LERR))
	return l.l.Output(2, fmt.Sprintln(msg...))
}

func (l *Logger) Fatal(msg ...interface{}) {
	l.l.SetPrefix(fmt.Sprintf("%s ", LFATAL))
	l.l.Output(2, fmt.Sprintln(msg...))
	os.Exit(1)
}
