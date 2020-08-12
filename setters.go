package nilog

import (
	"io"
	"log"
)

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
