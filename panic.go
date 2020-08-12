package nilog

import "fmt"

// Panic is equivalent to l.Print() followed by a call to panic().
func (l *Logger) Panic(v ...interface{}) {
	if l == nil {
		return
	}

	s := fmt.Sprint(v...)
	l.Output(2, s)
	panic(s)
}

// Panicf is equivalent to l.Printf() followed by a call to panic().
func (l *Logger) Panicf(format string, v ...interface{}) {
	if l == nil {
		return
	}

	s := fmt.Sprintf(format, v...)
	l.Output(2, s)
	panic(s)
}

// Panicln is equivalent to l.Println() followed by a call to panic().
func (l *Logger) Panicln(v ...interface{}) {
	if l == nil {
		return
	}

	s := fmt.Sprintln(v...)
	l.Output(2, s)
	panic(s)
}
