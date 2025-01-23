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
		{
			input: "at .LT ( http://localhost:8080/dist/js/chunk-M6AOQWLO.js:12217:35 )",
			want: &StackFrame{
				Function: ".LT",
				Url:      "http://localhost:8080/dist/js/chunk-M6AOQWLO.js",
				Line:     12217,
				Column:   35,
			},
		},
		{
			input: "at .LT ( http://localhost:8080/dist/js/chunk-M6AOQWLO.js:12217 )",
			want: &StackFrame{
				Function: ".LT",
				Url:      "http://localhost:8080/dist/js/chunk-M6AOQWLO.js",
				Line:     12217,
				Column:   -1,
			},
		},
		{
			input: "Error",
			want:  nil,
		},
		{
			input: "Error: something went wrong",
			want:  nil,
		},
		{
			input: "   at <anonymous>",
			want:  &StackFrame{Function: "<anonymous>", Url: "", Line: -1, Column: -1},
		},
		{
			input: "   at http://localhost:8080/dist/js/chunk-M6AOQWLO.js:12217:35",
			want:  &StackFrame{Function: "", Url: "http://localhost:8080/dist/js/chunk-M6AOQWLO.js", Line: 12217, Column: 35},
		},
		{
			input: "@http://path/to/file.js:48",
			want:  &StackFrame{Function: "", Url: "http://path/to/file.js", Line: 48, Column: -1},
		},
		{
			input: "printStackTrace()@file:///G:/js/stacktrace.js:18",
			want:  &StackFrame{Function: "printStackTrace()", Url: "file:///G:/js/stacktrace.js", Line: 18, Column: -1},
		},
		{
			input: ".plugin/e.fn[c]/<@http://path/to/file.js:1:1",
			want:  &StackFrame{Function: ".plugin/e.fn[c]/<", Url: "http://path/to/file.js", Line: 1, Column: 1},
		},
	}

	for _, tt := range tests {
		got := StackFrameFromString(tt.input)
		if got != nil || tt.want != nil {
			if got == nil {
				t.Fatalf("StackFrameFromString(%q) = nil, want %v", tt.input, tt.want)
			} else if tt.want == nil {
				t.Fatalf("StackFrameFromString(%q) = %v, want nil", tt.input, got)
			} else if *got != *tt.want {
				t.Errorf("StackFrameFromString(%q) = %v, want %v", tt.input, got, tt.want)
			}
		}
	}
}
