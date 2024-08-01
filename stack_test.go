package jsstacktrace

import "testing"

func TestStack(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{
			`    at EditLink (http://localhost:8080/dist/js/chunk-XJD7LI46.js:359:9)
    at renderWithHooks (http://localhost:8080/dist/js/chunk-M6AOQWLO.js:12217:26)
    at mountIndeterminateComponent (http://localhost:8080/dist/js/chunk-M6AOQWLO.js:15595:21)
    at beginWork (http://localhost:8080/dist/js/chunk-M6AOQWLO.js:16583:22)
    at beginWork$1 (http://localhost:8080/dist/js/chunk-M6AOQWLO.js:20422:22)
    at performUnitOfWork (http://localhost:8080/dist/js/chunk-M6AOQWLO.js:19867:21)
    at workLoopSync (http://localhost:8080/dist/js/chunk-M6AOQWLO.js:19806:13)
    at renderRootSync (http://localhost:8080/dist/js/chunk-M6AOQWLO.js:19785:15)
    at recoverFromConcurrentError (http://localhost:8080/dist/js/chunk-M6AOQWLO.js:19405:28)
    at performSyncWorkOnRoot (http://localhost:8080/dist/js/chunk-M6AOQWLO.js:19548:28)`,
			`at EditLink (http://localhost:8080/dist/js/chunk-XJD7LI46.js:359:9)
at renderWithHooks (http://localhost:8080/dist/js/chunk-M6AOQWLO.js:12217:26)
at mountIndeterminateComponent (http://localhost:8080/dist/js/chunk-M6AOQWLO.js:15595:21)
at beginWork (http://localhost:8080/dist/js/chunk-M6AOQWLO.js:16583:22)
at beginWork$1 (http://localhost:8080/dist/js/chunk-M6AOQWLO.js:20422:22)
at performUnitOfWork (http://localhost:8080/dist/js/chunk-M6AOQWLO.js:19867:21)
at workLoopSync (http://localhost:8080/dist/js/chunk-M6AOQWLO.js:19806:13)
at renderRootSync (http://localhost:8080/dist/js/chunk-M6AOQWLO.js:19785:15)
at recoverFromConcurrentError (http://localhost:8080/dist/js/chunk-M6AOQWLO.js:19405:28)
at performSyncWorkOnRoot (http://localhost:8080/dist/js/chunk-M6AOQWLO.js:19548:28)`,
		},
		{
			`EditLink@http://localhost:8080/dist/js/chunk-XJD7LI46.js:359:9
renderWithHooks@http://localhost:8080/dist/js/chunk-M6AOQWLO.js:12217:26
mountIndeterminateComponent@http://localhost:8080/dist/js/chunk-M6AOQWLO.js:15595:21
beginWork@http://localhost:8080/dist/js/chunk-M6AOQWLO.js:16583:22
beginWork$1@http://localhost:8080/dist/js/chunk-M6AOQWLO.js:20422:22
performUnitOfWork@http://localhost:8080/dist/js/chunk-M6AOQWLO.js:19870:21
workLoopSync@http://localhost:8080/dist/js/chunk-M6AOQWLO.js:19806:30
renderRootSync@http://localhost:8080/dist/js/chunk-M6AOQWLO.js:19785:15
recoverFromConcurrentError@http://localhost:8080/dist/js/chunk-M6AOQWLO.js:19405:42
performSyncWorkOnRoot@http://localhost:8080/dist/js/chunk-M6AOQWLO.js:19548:28`,
			`at EditLink (http://localhost:8080/dist/js/chunk-XJD7LI46.js:359:9)
at renderWithHooks (http://localhost:8080/dist/js/chunk-M6AOQWLO.js:12217:26)
at mountIndeterminateComponent (http://localhost:8080/dist/js/chunk-M6AOQWLO.js:15595:21)
at beginWork (http://localhost:8080/dist/js/chunk-M6AOQWLO.js:16583:22)
at beginWork$1 (http://localhost:8080/dist/js/chunk-M6AOQWLO.js:20422:22)
at performUnitOfWork (http://localhost:8080/dist/js/chunk-M6AOQWLO.js:19870:21)
at workLoopSync (http://localhost:8080/dist/js/chunk-M6AOQWLO.js:19806:30)
at renderRootSync (http://localhost:8080/dist/js/chunk-M6AOQWLO.js:19785:15)
at recoverFromConcurrentError (http://localhost:8080/dist/js/chunk-M6AOQWLO.js:19405:42)
at performSyncWorkOnRoot (http://localhost:8080/dist/js/chunk-M6AOQWLO.js:19548:28)`,
		},
	}

	for _, tt := range tests {
		stack := ParseStackTrace(tt.input)
		if stack.String() != tt.output {
			t.Errorf("ParseStackTrace failed: got:\n%s\nwant:\n%s", stack.String(), tt.output)
		}
	}
}
