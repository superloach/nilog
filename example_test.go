// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package nilog_test

import (
	"bytes"
	"fmt"

	"github.com/superloach/nilog"
)

func ExampleLogger() {
	var (
		buf    bytes.Buffer
		logger = nilog.New(&buf, "logger: ", nilog.Lshortfile)
	)

	logger.Print("Hello, log file!")

	fmt.Print(&buf)
	// Output:
	// logger: example_test.go:20: Hello, log file!
}

func ExampleLogger_Output() {
	var (
		buf    bytes.Buffer
		logger = nilog.New(&buf, "INFO: ", nilog.Lshortfile)

		infof = func(info string) {
			err := logger.Output(2, info)
			if err != nil {
				panic(err)
			}
		}
	)

	infof("Hello world")

	fmt.Print(&buf)
	// Output:
	// INFO: example_test.go:40: Hello world
}
