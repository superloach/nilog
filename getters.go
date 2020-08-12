package nilog

import (
	"io"
	"log"
)

// Flags wraps (*log.Logger)(l).Flags, returning -1 if l is nil.
func (l *Logger) Flags() int {
	if l == nil {
		return -1
	}

	return (*log.Logger)(l).Flags()
}

// Prefix wraps (*log.Logger)(l).Prefix, returning "" if l is nil.
func (l *Logger) Prefix() string {
	if l == nil {
		return ""
	}

	return (*log.Logger)(l).Prefix()
}

// Writer wraps (*log.Logger)(l).Writer, returning nil if l is nil.
func (l *Logger) Writer() io.Writer {
	if l == nil {
		return nil
	}

	return (*log.Logger)(l).Writer()
}
