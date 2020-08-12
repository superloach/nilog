package nilog

import "fmt"

// Print calls l.Output to print to the logger.
// Arguments are handled in the manner of fmt.Print.
func (l *Logger) Print(v ...interface{}) {
	if l == nil {
		return
	}

	l.Output(2, fmt.Sprint(v...))
}

// Printf calls l.Output to print to the logger.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Printf(format string, v ...interface{}) {
	if l == nil {
		return
	}

	l.Output(2, fmt.Sprintf(format, v...))
}

// Println calls l.Output to print to the logger.
// Arguments are handled in the manner of fmt.Println.
func (l *Logger) Println(v ...interface{}) {
	if l == nil {
		return
	}

	l.Output(2, fmt.Sprintln(v...))
}
