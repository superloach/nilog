package nilog

import (
	"io"
	"log"
)

// Logger wraps *log.Logger, making a nil value a no-op.
type Logger log.Logger

// New wraps log.New into a Logger.
func New(out io.Writer, prefix string, flag int) *Logger {
	return (*Logger)(log.New(out, prefix, flag))
}

// Fatal wraps (*l).Fatal, returning early if l is nil.
func (l *Logger) Fatal(v ...interface{}) {
	if l == nil {
		return
	}

	(*l).Fatal(v...)
}

// Fatalf wraps (*l).Fatalf, returning early if l is nil.
func (l *Logger) Fatalf(format string, v ...interface{}) {
	if l == nil {
		return
	}

	(*l).Fatalf(format, v...)
}

// Fatalln wraps (*l).Fatalln, returning early if l is nil.
func (l *Logger) Fatalln(v ...interface{}) {
	if l == nil {
		return
	}

	(*l).Fatalln(v...)
}

// Flags wraps (*l).Flags, returning early if l is nil.
func (l *Logger) Flags() int {
	if l == nil {
		return -1
	}

	return (*l).Flags()
}

// Output wraps (*l).Output, returning early if l is nil.
func (l *Logger) Output(calldepth int, s string) error {
	if l == nil {
		return nil
	}

	return (*l).Output(calldepth, s)
}

// Panic wraps (*l).Panic, returning early if l is nil.
func (l *Logger) Panic(v ...interface{}) {
	if l == nil {
		return
	}

	(*l).Panic(v...)
}

// Panicf wraps (*l).Panicf, returning early if l is nil.
func (l *Logger) Panicf(format string, v ...interface{}) {
	if l == nil {
		return
	}

	(*l).Panicf(format, v...)
}

// Panicln wraps (*l).Panicln, returning early if l is nil.
func (l *Logger) Panicln(v ...interface{}) {
	if l == nil {
		return
	}

	(*l).Panicln(v...)
}

// Prefix wraps (*l).Prefix, returning early if l is nil.
func (l *Logger) Prefix() string {
	if l == nil {
		return ""
	}

	return (*l).Prefix()
}

// Print wraps (*l).Print, returning early if l is nil.
func (l *Logger) Print(v ...interface{}) {
	if l == nil {
		return
	}

	(*l).Print(v...)
}

// Printf wraps (*l).Printf, returning early if l is nil.
func (l *Logger) Printf(format string, v ...interface{}) {
	if l == nil {
		return
	}

	(*l).Printf(format, v...)
}

// Println wraps (*l).Println, returning early if l is nil.
func (l *Logger) Println(v ...interface{}) {
	if l == nil {
		return
	}

	(*l).Println(v...)
}

// SetFlags wraps (*l).SetFlags, returning early if l is nil.
func (l *Logger) SetFlags(flag int) {
	if l == nil {
		return
	}

	(*l).SetFlags(flag)
}

// SetOutput wraps (*l).SetOutput, returning early if l is nil.
func (l *Logger) SetOutput(w io.Writer) {
	if l == nil {
		return
	}

	(*l).SetOutput(w)
}

// SetPrefix wraps (*l).SetPrefix, returning early if l is nil.
func (l *Logger) SetPrefix(prefix string) {
	if l == nil {
		return
	}

	(*l).SetPrefix(prefix)
}

// Writer wraps (*l).Writer, returning early if l is nil.
func (l *Logger) Writer() io.Writer {
	if l == nil {
		return nil
	}

	return (*l).Writer()
}
