package nilog

import (
	"fmt"
	"os"
)

// Fatal is equivalent to l.Print() followed by a call to os.Exit(1).
func (l *Logger) Fatal(v ...interface{}) {
	if l == nil {
		return
	}

	l.Output(2, fmt.Sprint(v...))
	os.Exit(1)
}

// Fatalf is equivalent to l.Printf() followed by a call to os.Exit(1).
func (l *Logger) Fatalf(format string, v ...interface{}) {
	if l == nil {
		return
	}

	l.Output(2, fmt.Sprintf(format, v...))
	os.Exit(1)
}

// Fatalln is equivalent to l.Println() followed by a call to os.Exit(1).
func (l *Logger) Fatalln(v ...interface{}) {
	if l == nil {
		return
	}

	l.Output(2, fmt.Sprintln(v...))
	os.Exit(1)
}
