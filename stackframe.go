package jsstacktrace

import (
	"regexp"
	"strconv"
)

type StackFrame struct {
	File     string
	Line     int
	Column   int
	Function string
}

func (sf *StackFrame) String() string {
	return sf.Function + " @ " + sf.File + ":" + strconv.Itoa(sf.Line) + ":" + strconv.Itoa(sf.Column)
}

// Stack trace formats:
// Safari and Firefox: "function@url:line:column"
// Chrome and Edge: "\s+at function (url:line:column)"

var stackFrameSafari = regexp.MustCompile(`(?P<function>[a-zA-Z_$][a-zA-Z0-9_$]*)@(?P<file>.*):(?P<line>\d+):(?P<column>\d+)`)
var stackFrameChrome = regexp.MustCompile(`\s+at\s+(?P<function>[a-zA-Z_$][a-zA-Z0-9_$]*) \((?P<file>.*):(?P<line>\d+):(?P<column>\d+)\)`)

func StackFrameFromString(s string) *StackFrame {
	// check which format the stack trace has
	var matches []string
	if stackFrameSafari.MatchString(s) {
		matches = stackFrameSafari.FindStringSubmatch(s)
	} else if stackFrameChrome.MatchString(s) {
		matches = stackFrameChrome.FindStringSubmatch(s)
	} else {
		return nil
	}
	line, _ := strconv.Atoi(matches[3])
	column, _ := strconv.Atoi(matches[4])
	return &StackFrame{Function: matches[1], File: matches[2], Line: line, Column: column}
}
