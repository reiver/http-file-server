package main

import (
	"fmt"
	"os"
)

func log(a ...interface{}) {
	s := fmt.Sprint(a...)

	logf("%s", s)
}

func logf(format string, a ...interface{}) {
	format = "[LOG] " + format + "\n"

	fmt.Fprintf(os.Stdout, format, a...)
}

func logerror(a ...interface{}) {
	s := fmt.Sprint(a...)

	logerrorf("%s", s)
}

func logerrorf(format string, a ...interface{}) {
	format = "[ERROR] " + format + "\n"

	fmt.Fprintf(os.Stderr, format, a...)
}
