package golog

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

type Log struct {
	sync.Mutex
	level int
	out   io.Writer
}

const (
	Linfo = 1 << iota
	Ldebug
	Lother
)

func NewLog(out io.Writer, level int) *Log {
	return &Log{out: out, level: level}
}
func (l *Log) SetLevel(level int) {
	l.Lock()
	defer l.Unlock()
	l.level = level
}

func (l *Log) format(now time.Time, file string, line int, flag int) (s string) {

	if flag&Linfo != 0 {
		s += fmt.Sprint(now.Year(), "/", int(now.Month()), "/", now.Day()) + " "
	}
	s += fmt.Sprint(now.Hour(), ":", now.Minute(), ":", now.Second())

	if flag&Linfo != 0 {
		return
	}
	short := file
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			short = file[i+1:]
			break
		}
	}
	s += " " + short + ":" + strconv.Itoa(line)
	return
}

func (l *Log) Output(prefix, s string, flag int) {
	now := time.Now()

	l.Lock()
	defer l.Unlock()

	_, file, line, _ := runtime.Caller(2)
	formatStr := l.format(now, file, line, flag)
	outstring := prefix + " " + formatStr + " " + s
	l.out.Write([]byte(outstring))
}

func (l *Log) Info(args ...interface{}) {
	if l.level&Linfo != 0 {
		l.Output("[Info]", fmt.Sprintln(args...), Linfo)
	}
}

func (l *Log) Debug(args ...interface{}) {
	if l.level&Ldebug != 0 {
		l.Output("[Debug]", fmt.Sprintln(args...), Ldebug)
	}
}

func (l *Log) Panic(args ...interface{}) {
	s := fmt.Sprintln(args...)
	l.Output("[Panic]", s, Lother)
	panic(s)
}

func (l *Log) Fatal(args ...interface{}) {
	l.Output("[Fatal]", fmt.Sprintln(args...), Lother)
	os.Exit(1)
}

func (l *Log) Println(prefix string, args ...interface{}) {
	s := fmt.Sprintln(args...)
	l.Output(prefix, s, Lother)
}

func (l *Log) Infof(format string, args ...interface{}) {
	if l.level&Linfo != 0 {
		l.Output("[Info]", fmt.Sprintf(format, args...), Linfo)
	}
}

func (l *Log) Debugf(format string, args ...interface{}) {
	if l.level&Ldebug != 0 {
		l.Output("[Debug]", fmt.Sprintf(format, args...), Ldebug)
	}
}

func (l *Log) Panicf(format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	l.Output("[Panic]", s, Lother)
	panic(s)
}

func (l *Log) Printf(prefix string, format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	l.Output(prefix, s, Lother)
}

func (l *Log) Fatalf(format string, args ...interface{}) {
	l.Output("[Fatal]", fmt.Sprintf(format, args...), Lother)
	os.Exit(1)
}
