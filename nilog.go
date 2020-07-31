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

// Fatal wraps (*log.Logger)(l).Fatal, returning early if l is nil.
func (l *Logger) Fatal(v ...interface{}) {
	if l == nil {
		return
	}

	(*log.Logger)(l).Fatal(v...)
}

// Fatalf wraps (*log.Logger)(l).Fatalf, returning early if l is nil.
func (l *Logger) Fatalf(format string, v ...interface{}) {
	if l == nil {
		return
	}

	(*log.Logger)(l).Fatalf(format, v...)
}

// Fatalln wraps (*log.Logger)(l).Fatalln, returning early if l is nil.
func (l *Logger) Fatalln(v ...interface{}) {
	if l == nil {
		return
	}

	(*log.Logger)(l).Fatalln(v...)
}

// Flags wraps (*log.Logger)(l).Flags, returning early if l is nil.
func (l *Logger) Flags() int {
	if l == nil {
		return -1
	}

	return (*log.Logger)(l).Flags()
}

// Output wraps (*log.Logger)(l).Output, returning early if l is nil.
func (l *Logger) Output(calldepth int, s string) error {
	if l == nil {
		return nil
	}

	return (*log.Logger)(l).Output(calldepth, s)
}

// Panic wraps (*log.Logger)(l).Panic, returning early if l is nil.
func (l *Logger) Panic(v ...interface{}) {
	if l == nil {
		return
	}

	(*log.Logger)(l).Panic(v...)
}

// Panicf wraps (*log.Logger)(l).Panicf, returning early if l is nil.
func (l *Logger) Panicf(format string, v ...interface{}) {
	if l == nil {
		return
	}

	(*log.Logger)(l).Panicf(format, v...)
}

// Panicln wraps (*log.Logger)(l).Panicln, returning early if l is nil.
func (l *Logger) Panicln(v ...interface{}) {
	if l == nil {
		return
	}

	(*log.Logger)(l).Panicln(v...)
}

// Prefix wraps (*log.Logger)(l).Prefix, returning early if l is nil.
func (l *Logger) Prefix() string {
	if l == nil {
		return ""
	}

	return (*log.Logger)(l).Prefix()
}

// Print wraps (*log.Logger)(l).Print, returning early if l is nil.
func (l *Logger) Print(v ...interface{}) {
	if l == nil {
		return
	}

	(*log.Logger)(l).Print(v...)
}

// Printf wraps (*log.Logger)(l).Printf, returning early if l is nil.
func (l *Logger) Printf(format string, v ...interface{}) {
	if l == nil {
		return
	}

	(*log.Logger)(l).Printf(format, v...)
}

// Println wraps (*log.Logger)(l).Println, returning early if l is nil.
func (l *Logger) Println(v ...interface{}) {
	if l == nil {
		return
	}

	(*log.Logger)(l).Println(v...)
}

// SetFlags wraps (*log.Logger)(l).SetFlags, returning early if l is nil.
func (l *Logger) SetFlags(flag int) {
	if l == nil {
		return
	}

	(*log.Logger)(l).SetFlags(flag)
}

// SetOutput wraps (*log.Logger)(l).SetOutput, returning early if l is nil.
func (l *Logger) SetOutput(w io.Writer) {
	if l == nil {
		return
	}

	(*log.Logger)(l).SetOutput(w)
}

// SetPrefix wraps (*log.Logger)(l).SetPrefix, returning early if l is nil.
func (l *Logger) SetPrefix(prefix string) {
	if l == nil {
		return
	}

	(*log.Logger)(l).SetPrefix(prefix)
}

// Writer wraps (*log.Logger)(l).Writer, returning early if l is nil.
func (l *Logger) Writer() io.Writer {
	if l == nil {
		return nil
	}

	return (*log.Logger)(l).Writer()
}
