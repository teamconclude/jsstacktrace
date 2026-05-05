package jsstacktrace

import (
	"bytes"
	"compress/gzip"
	"io"
	"net/url"
	"os"

	"github.com/go-sourcemap/sourcemap"
)

type JSMap struct {
	baseDir  string
	mapFiles map[string]*sourcemap.Consumer
}

func NewJSMap(baseDir string) *JSMap {
	return &JSMap{
		baseDir:  baseDir,
		mapFiles: make(map[string]*sourcemap.Consumer),
	}
}

func (m *JSMap) getMapFile(mapURL string) (*sourcemap.Consumer, error) {
	// remove the protocol and domain from the URL and prefix it with the baseDir
	// to make it relative to the baseDir

	u, err := url.Parse(mapURL)
	if err != nil {
		return nil, err
	}

	if c, ok := m.mapFiles[u.Path]; ok {
		return c, nil
	}

	filename := m.baseDir + u.Path + ".map"

	data, err := os.ReadFile(filename)
	if err != nil {
		// Production builds may pre-gzip source maps and delete the original.
		gzData, gzErr := os.ReadFile(filename + ".gz")
		if gzErr != nil {
			return nil, err
		}
		zr, zerr := gzip.NewReader(bytes.NewReader(gzData))
		if zerr != nil {
			return nil, zerr
		}
		defer zr.Close()
		data, err = io.ReadAll(zr)
		if err != nil {
			return nil, err
		}
	}

	smap, err := sourcemap.Parse(mapURL, data)
	if err != nil {
		return nil, err
	}

	m.mapFiles[u.Path] = smap
	return smap, nil
}

func (m *JSMap) ConvertFrame(stackFrame StackFrame) StackFrame {
	if stackFrame.Other != "" {
		return stackFrame
	}
	if stackFrame.Url == "" {
		return stackFrame
	}

	mapFile, err := m.getMapFile(stackFrame.Url)
	if err != nil {
		return stackFrame
	}

	file, fn, line, column, ok := mapFile.Source(stackFrame.Line, stackFrame.Column)
	if !ok {
		return stackFrame
	}
	if fn == "" {
		fn = stackFrame.Function
	}
	return StackFrame{
		Url:      file,
		Line:     line,
		Column:   column,
		Function: fn,
	}

}

func (m *JSMap) ConvertStackTrace(stackTrace StackTrace) StackTrace {
	var convertedStackTrace StackTrace
	for _, frame := range stackTrace {
		convertedStackTrace = append(convertedStackTrace, m.ConvertFrame(frame))
	}
	return convertedStackTrace
}

func (m *JSMap) ConvertStackTraceString(stackTrace string) string {
	return m.ConvertStackTrace(ParseStackTrace(stackTrace)).String()
}
