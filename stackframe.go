package jsstacktrace

import (
	"regexp"
	"strconv"
	"strings"
)

type StackFrame struct {
	Url      string
	Line     int
	Column   int
	Function string
}

func (sf *StackFrame) String() string {
	result := "at " + sf.Function + " (" + sf.Url
	if sf.Line >= 0 {
		result += ":" + strconv.Itoa(sf.Line)
		if sf.Column >= 0 {
			result += ":" + strconv.Itoa(sf.Column)
		}
	}
	return result + ")"
}

// Stack trace formats:
// Safari and Firefox: "function@url:line:column"
// Chrome and Edge: "\s+at function (url:line:column)"

var chromeFormat = regexp.MustCompile(`\s*at\s+(.*)`)
var safariFormat = regexp.MustCompile(`\s*(.*)@(.*)`)

var stackFrameChrome = regexp.MustCompile(`([^( ]+)\s+\(\s*(?P<file>[^\s]+)\s*\)\s*`)

var splitLocationWithColumn = regexp.MustCompile(`(?P<url>.+):(?P<line>\d+):(?P<column>\d+)\s*`)
var splitLocationWithoutColumn = regexp.MustCompile(`(?P<url>.+):(?P<line>\d+)\s*`)

func parseLocation(location string) (string, int, int) {
	url := location
	line := -1
	column := -1

	if splitLocationWithColumn.MatchString(location) {
		matches := splitLocationWithColumn.FindStringSubmatch(location)
		url = matches[1]
		line, _ = strconv.Atoi(matches[2])
		column, _ = strconv.Atoi(matches[3])
	} else if splitLocationWithoutColumn.MatchString(location) {
		matches := splitLocationWithoutColumn.FindStringSubmatch(location)
		url = matches[1]
		line, _ = strconv.Atoi(matches[2])
	}

	return url, line, column
}

func parseFrameChrome(s string) *StackFrame {
	// log.Default().Println("parseFrameChrome: ", s)
	matches := stackFrameChrome.FindStringSubmatch(s)
	loc := s
	function := ""
	if len(matches) == 0 {
		if strings.Contains(s, ":") {
			loc = s
		} else {
			function = s
			loc = ""
		}
	} else {
		function = matches[1]
		loc = matches[2]
	}

	url, line, column := parseLocation(loc)

	return &StackFrame{Function: function, Url: url, Line: line, Column: column}
}

func parseFrameSafari(f string, l string) *StackFrame {
	function := f
	url, line, column := parseLocation(l)

	return &StackFrame{Function: function, Url: url, Line: line, Column: column}
}

func StackFrameFromString(s string) *StackFrame {
	// check which format the stack trace has
	// log.Default().Println("StackFrameFromString: ", s)
	if chromeFormat.MatchString(s) {
		// log.Default().Println("StackFrameFromString: chromeFormat")
		matches := chromeFormat.FindStringSubmatch(s)
		return parseFrameChrome(matches[1])
	} else if safariFormat.MatchString(s) {
		// log.Default().Println("StackFrameFromString: safariFormat")
		matches := safariFormat.FindStringSubmatch(s)
		return parseFrameSafari(matches[1], matches[2])
	} else {
		return nil
	}
}
