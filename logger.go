package logger

import (
	"io"
	"log"
)

type logger struct {
	out io.Writer
}

type Logger interface {
	Write(...interface{})
}

func New(w io.Writer) Logger {
	return &logger{out: w}
}

func (l *logger) Write(a ...interface{}) {
	log.SetOutput(l.out)
	log.Println(a...)
}
