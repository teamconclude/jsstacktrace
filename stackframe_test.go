package jsstacktrace

import (
	"testing"
)

func TestStackFrameFromString(t *testing.T) {
	tests := []struct {
		input string
		want  *StackFrame
	}{
		{
			input: "renderWithHooks@http://localhost:8080/dist/js/chunk-M6AOQWLO.js:12217:35",
			want: &StackFrame{
				Function: "renderWithHooks",
				Url:      "http://localhost:8080/dist/js/chunk-M6AOQWLO.js",
				Line:     12217,
				Column:   35,
			},
		},
		{
			input: "    at renderWithHooks (http://localhost:8080/dist/js/chunk-M6AOQWLO.js:12217:35)",
			want: &StackFrame{
				Function: "renderWithHooks",
				Url:      "http://localhost:8080/dist/js/chunk-M6AOQWLO.js",
				Line:     12217,
				Column:   35,
			},
		},
	}

	for _, tt := range tests {
		got := StackFrameFromString(tt.input)
		if got == nil {
			t.Fatalf("StackFrameFromString(%q) = nil, want %v", tt.input, tt.want)
		}
		if *got != *tt.want {
			t.Errorf("StackFrameFromString(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}
}
