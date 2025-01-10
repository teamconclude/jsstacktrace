package jsstacktrace

import (
	"log"
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
