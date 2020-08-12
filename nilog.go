// Package nilog provides a wrapper for log.Logger, making a nil value a no-op.
package nilog

import (
	"io"
	"log"
)

// Logger wraps log.Logger, making a nil value a no-op.
type Logger log.Logger

// New wraps log.New into a *Logger.
func New(out io.Writer, prefix string, flag int) *Logger {
	return (*Logger)(log.New(out, prefix, flag))
}

// Output wraps (*log.Logger)(l).Output, returning early if l is nil, and incrementing the calldepth by 1.
func (l *Logger) Output(calldepth int, s string) error {
	if l == nil {
		return nil
	}

	return (*log.Logger)(l).Output(calldepth+1, s)
}
