package jsstacktrace

import "strings"

type Stack []StackFrame

func ParseStackTrace(s string) Stack {
	lines := strings.Split(s, "\n")

	stack := make(Stack, 0, len(lines))
	for _, line := range lines {
		if sf := StackFrameFromString(line); sf != nil {
			stack = append(stack, *sf)
		}
	}

	return stack
}

func (s Stack) String() string {
	var b strings.Builder
	for i, sf := range s {
		b.WriteString(sf.String())
		if i < len(s)-1 {
			b.WriteString("\n")
		}
	}
	return b.String()
}
