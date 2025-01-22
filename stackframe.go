package jsstacktrace

import (
	"regexp"
	"strconv"
)

type StackFrame struct {
	Url      string
	Line     int
	Column   int
	Function string
}

func (sf *StackFrame) String() string {
	return "at " + sf.Function + " (" + sf.Url + ":" + strconv.Itoa(sf.Line) + ":" + strconv.Itoa(sf.Column) + ")"
}

// Stack trace formats:
// Safari and Firefox: "function@url:line:column"
// Chrome and Edge: "\s+at function (url:line:column)"

var stackFrameChrome = regexp.MustCompile(`\s*at\s+(?P<function>[^\s]+)\s+\(\s*(?P<file>[^\s]+)\s*\)\s*`)
var stackFrameSafari = regexp.MustCompile(`\s*(?P<function>[^@\s]+)@(?P<file>[^\s]+)\s*`)

var splitLocationWithColumn = regexp.MustCompile(`(?P<url>.+):(?P<line>\d+):(?P<column>\d+)`)
var splitLocationWithoutColumn = regexp.MustCompile(`(?P<url>.+):(?P<line>\d+)`)

func StackFrameFromString(s string) *StackFrame {
	// check which format the stack trace has
	var matches []string
	if stackFrameSafari.MatchString(s) {
		matches = stackFrameSafari.FindStringSubmatch(s)
	} else if stackFrameChrome.MatchString(s) {
		matches = stackFrameChrome.FindStringSubmatch(s)
	} else {
		return &StackFrame{Function: "JSStackframe: not converted: ", Url: s, Line: 0, Column: 0}
	}
	function := matches[1]
	location := matches[2]
	var url string
	var line int
	var column int

	var err error
	if splitLocationWithColumn.MatchString(location) {
		matches = splitLocationWithColumn.FindStringSubmatch(location)
		url = matches[1]
		column, err = strconv.Atoi(matches[3])
		if err == nil {
			line, err = strconv.Atoi(matches[2])
		}
	}
	if (url == "" || err != nil) && splitLocationWithoutColumn.MatchString(location) {
		matches = splitLocationWithoutColumn.FindStringSubmatch(location)
		url = matches[1]
		line, err = strconv.Atoi(matches[2])
		column = 0
	}
	if url == "" || err != nil {
		url = location
	}

	return &StackFrame{Function: function, Url: url, Line: line, Column: column}
}
