package jsstacktrace

import (
	"bytes"
	"compress/gzip"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestMap(t *testing.T) {

	frame := StackFrameFromString("    at EditLink (http://localhost:8080/dist/js/chunk-XJD7LI46.js:359:9)")
	if frame == nil {
		t.Errorf("StackFrameFromString failed")
		return
	}

	jsmap := NewJSMap("testdata")

	convertedFrame := jsmap.ConvertFrame(*frame)

	log.Printf("%+v", convertedFrame)

	expected := StackFrame{
		Function: "EditLink",
		Url:      "http://localhost:8080/js/modules/link/link/ChannelSettings.tsx",
		Line:     483,
		Column:   10,
	}

	if convertedFrame.Url != expected.Url || convertedFrame.Line != expected.Line || convertedFrame.Column != expected.Column || convertedFrame.Function != expected.Function {
		t.Errorf("ConvertFrame failed")
	}
}

// TestMapGzipped verifies that getMapFile falls back to <file>.map.gz when
// the uncompressed sibling has been removed (production Vite builds do this).
func TestMapGzipped(t *testing.T) {
	src, err := os.ReadFile("testdata/dist/js/chunk-XJD7LI46.js.map")
	if err != nil {
		t.Fatalf("read fixture: %v", err)
	}

	tmp := t.TempDir()
	dstDir := filepath.Join(tmp, "dist", "js")
	if err := os.MkdirAll(dstDir, 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}

	var buf bytes.Buffer
	gw := gzip.NewWriter(&buf)
	if _, err := gw.Write(src); err != nil {
		t.Fatalf("gzip write: %v", err)
	}
	if err := gw.Close(); err != nil {
		t.Fatalf("gzip close: %v", err)
	}
	if err := os.WriteFile(filepath.Join(dstDir, "chunk-XJD7LI46.js.map.gz"), buf.Bytes(), 0o644); err != nil {
		t.Fatalf("write gz: %v", err)
	}

	frame := StackFrameFromString("    at EditLink (http://localhost:8080/dist/js/chunk-XJD7LI46.js:359:9)")
	if frame == nil {
		t.Fatal("StackFrameFromString failed")
	}

	jsmap := NewJSMap(tmp)
	got := jsmap.ConvertFrame(*frame)

	expected := StackFrame{
		Function: "EditLink",
		Url:      "http://localhost:8080/js/modules/link/link/ChannelSettings.tsx",
		Line:     483,
		Column:   10,
	}
	if got.Url != expected.Url || got.Line != expected.Line || got.Column != expected.Column || got.Function != expected.Function {
		t.Errorf("ConvertFrame from .map.gz failed: got %+v, want %+v", got, expected)
	}
}
