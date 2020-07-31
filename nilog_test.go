// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package nilog_test

// These tests are too simple.

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strings"
	"testing"
	"time"

	"log"

	. "github.com/superloach/nilog"
)

func itoa(buf *[]byte, i int, wid int) {
	// Assemble decimal in reverse order.
	var b [20]byte
	bp := len(b) - 1
	for i >= 10 || wid > 1 {
		wid--
		q := i / 10
		b[bp] = byte('0' + i - q*10)
		bp--
		i = q
	}
	// i < 10
	b[bp] = byte('0' + i)
	*buf = append(*buf, b[bp:]...)
}

const (
	Rdate         = `[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9]`
	Rtime         = `[0-9][0-9]:[0-9][0-9]:[0-9][0-9]`
	Rmicroseconds = `\.[0-9][0-9][0-9][0-9][0-9][0-9]`
	Rline         = `(80|82):` // must update if the calls to l.Printf / l.Print below move
	Rlongfile     = `.*/[A-Za-z0-9_\-]+\.go:` + Rline
	Rshortfile    = `[A-Za-z0-9_\-]+\.go:` + Rline
)

type tester struct {
	flag    int
	prefix  string
	pattern string // regexp that log output must match; we add ^ and expected_text$ always
}

var tests = []tester{
	// individual pieces:
	{0, "", ""},
	{0, "XXX", "XXX"},
	{log.Ldate, "", Rdate + " "},
	{log.Ltime, "", Rtime + " "},
	{log.Ltime | log.Lmsgprefix, "XXX", Rtime + " XXX"},
	{log.Ltime | log.Lmicroseconds, "", Rtime + Rmicroseconds + " "},
	{log.Lmicroseconds, "", Rtime + Rmicroseconds + " "}, // microsec implies time
	{log.Llongfile, "", Rlongfile + " "},
	{log.Lshortfile, "", Rshortfile + " "},
	{log.Llongfile | log.Lshortfile, "", Rshortfile + " "}, // shortfile overrides longfile
	// everything at once:
	{log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile, "XXX", "XXX" + Rdate + " " + Rtime + Rmicroseconds + " " + Rlongfile + " "},
	{log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile, "XXX", "XXX" + Rdate + " " + Rtime + Rmicroseconds + " " + Rshortfile + " "},
	{log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile | log.Lmsgprefix, "XXX", Rdate + " " + Rtime + Rmicroseconds + " " + Rlongfile + " XXX"},
	{log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile | log.Lmsgprefix, "XXX", Rdate + " " + Rtime + Rmicroseconds + " " + Rshortfile + " XXX"},
}

// Test using Println("hello", 23, "world") or using Printf("hello %d world", 23)
func testPrint(t *testing.T, flag int, prefix string, pattern string, useFormat bool) {
	buf := new(bytes.Buffer)
	log.SetOutput(buf)
	log.SetFlags(flag)
	log.SetPrefix(prefix)
	if useFormat {
		log.Printf("hello %d world", 23)
	} else {
		log.Println("hello", 23, "world")
	}
	line := buf.String()
	line = line[0 : len(line)-1]
	pattern = "^" + pattern + "hello 23 world$"
	matched, err := regexp.MatchString(pattern, line)
	if err != nil {
		t.Fatal("pattern did not compile:", err)
	}
	if !matched {
		t.Errorf("log output should match %q is %q", pattern, line)
	}
	log.SetOutput(os.Stderr)
}

func TestAll(t *testing.T) {
	for _, testcase := range tests {
		testPrint(t, testcase.flag, testcase.prefix, testcase.pattern, false)
		testPrint(t, testcase.flag, testcase.prefix, testcase.pattern, true)
	}
}

func TestOutput(t *testing.T) {
	const testString = "test"
	var b bytes.Buffer
	l := New(&b, "", 0)
	l.Println(testString)
	if expect := testString + "\n"; b.String() != expect {
		t.Errorf("log output should match %q is %q", expect, b.String())
	}
}

func TestOutputRace(t *testing.T) {
	var b bytes.Buffer
	l := New(&b, "", 0)
	for i := 0; i < 100; i++ {
		go func() {
			l.SetFlags(0)
		}()
		l.Output(0, "")
	}
}

func TestFlagAndPrefixSetting(t *testing.T) {
	var b bytes.Buffer
	l := New(&b, "Test:", log.LstdFlags)
	f := l.Flags()
	if f != log.LstdFlags {
		t.Errorf("Flags 1: expected %x got %x", log.LstdFlags, f)
	}
	l.SetFlags(f | log.Lmicroseconds)
	f = l.Flags()
	if f != log.LstdFlags|log.Lmicroseconds {
		t.Errorf("Flags 2: expected %x got %x", log.LstdFlags|log.Lmicroseconds, f)
	}
	p := l.Prefix()
	if p != "Test:" {
		t.Errorf(`Prefix: expected "Test:" got %q`, p)
	}
	l.SetPrefix("Reality:")
	p = l.Prefix()
	if p != "Reality:" {
		t.Errorf(`Prefix: expected "Reality:" got %q`, p)
	}
	// Verify a log message looks right, with our prefix and microseconds present.
	l.Print("hello")
	pattern := "^Reality:" + Rdate + " " + Rtime + Rmicroseconds + " hello\n"
	matched, err := regexp.Match(pattern, b.Bytes())
	if err != nil {
		t.Fatalf("pattern %q did not compile: %s", pattern, err)
	}
	if !matched {
		t.Error("message did not match pattern")
	}
}

func TestUTCFlag(t *testing.T) {
	var b bytes.Buffer
	l := New(&b, "Test:", log.LstdFlags)
	l.SetFlags(log.Ldate | log.Ltime | log.LUTC)
	// Verify a log message looks right in the right time zone. Quantize to the second only.
	now := time.Now().UTC()
	l.Print("hello")
	want := fmt.Sprintf("Test:%d/%.2d/%.2d %.2d:%.2d:%.2d hello\n",
		now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	got := b.String()
	if got == want {
		return
	}
	// It's possible we crossed a second boundary between getting now and logging,
	// so add a second and try again. This should very nearly always work.
	now = now.Add(time.Second)
	want = fmt.Sprintf("Test:%d/%.2d/%.2d %.2d:%.2d:%.2d hello\n",
		now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	if got == want {
		return
	}
	t.Errorf("got %q; want %q", got, want)
}

func TestEmptyPrintCreatesLine(t *testing.T) {
	var b bytes.Buffer
	l := New(&b, "Header:", log.LstdFlags)
	l.Print()
	l.Println("non-empty")
	output := b.String()
	if n := strings.Count(output, "Header"); n != 2 {
		t.Errorf("expected 2 headers, got %d", n)
	}
	if n := strings.Count(output, "\n"); n != 2 {
		t.Errorf("expected 2 lines, got %d", n)
	}
}

func BenchmarkItoa(b *testing.B) {
	dst := make([]byte, 0, 64)
	for i := 0; i < b.N; i++ {
		dst = dst[0:0]
		itoa(&dst, 2015, 4)   // year
		itoa(&dst, 1, 2)      // month
		itoa(&dst, 30, 2)     // day
		itoa(&dst, 12, 2)     // hour
		itoa(&dst, 56, 2)     // minute
		itoa(&dst, 0, 2)      // second
		itoa(&dst, 987654, 6) // microsecond
	}
}

func BenchmarkPrintln(b *testing.B) {
	const testString = "test"
	var buf bytes.Buffer
	l := New(&buf, "", log.LstdFlags)
	for i := 0; i < b.N; i++ {
		buf.Reset()
		l.Println(testString)
	}
}

func BenchmarkPrintlnNoFlags(b *testing.B) {
	const testString = "test"
	var buf bytes.Buffer
	l := New(&buf, "", 0)
	for i := 0; i < b.N; i++ {
		buf.Reset()
		l.Println(testString)
	}
}
