package logd

import (
	"fmt"
	"io"
	"os"
	"sync"
)

type Logger struct {
	mu   sync.Mutex
	buf  []byte
	out  io.Writer
	fail func(msg ...interface{})
}

func NewLogger(out io.Writer, f func(...interface{})) *Logger {
	return &Logger{
		mu:   sync.Mutex{},
		buf:  make([]byte, 1<<10),
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

func (l *Logger) InfoPrintln(v ...interface{}) {
	l.mu.Lock()
	l.buf = append(l.buf, []byte("INFO: "+fmt.Sprint(v...)+"\n")...)
	l.mu.Unlock()
	l.write()
}

func (l *Logger) WarnPrintln(v ...interface{}) {
	l.mu.Lock()
	l.buf = append(l.buf, []byte("WARNING: "+fmt.Sprint(v...)+"\n")...)
	l.mu.Unlock()
	l.write()
}

func (l *Logger) ErrPrintln(v ...interface{}) {
	l.mu.Lock()
	l.buf = append(l.buf, []byte("ERROR: "+fmt.Sprint(v...)+"\n")...)
	l.mu.Unlock()
	l.write()
}

func (l *Logger) Fatal(v ...interface{}) {
	l.mu.Lock()
	l.buf = append(l.buf, []byte("FATAL: "+fmt.Sprint(v...)+"\n")...)
	l.mu.Unlock()
	l.write()
	os.Exit(1)
}

func (l *Logger) write() {
	l.mu.Lock()
	_, err := l.out.Write(l.buf)
	l.buf = l.buf[:0]
	l.mu.Unlock()
	if err != nil {
		l.fail(err)
	}
}
