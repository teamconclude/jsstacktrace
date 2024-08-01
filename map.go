package jsstacktrace

import (
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

	url, err := url.Parse(mapURL)
	if err != nil {
		return nil, err
	}

	path := url.Path

	if c, ok := m.mapFiles[url.Path]; ok {
		return c, nil
	}

	filename := m.baseDir + path + ".map"

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	smap, err := sourcemap.Parse(mapURL, data)
	if err != nil {
		return nil, err
	}

	m.mapFiles[mapURL] = smap
	return smap, nil
}

func (m *JSMap) ConvertFrame(stackFrame *StackFrame) StackFrame {
	if stackFrame.Url == "" {
		return *stackFrame
	}

	mapFile, err := m.getMapFile(stackFrame.Url)
	if err != nil {
		return *stackFrame
	}

	file, fn, line, column, ok := mapFile.Source(stackFrame.Line, stackFrame.Column)
	if !ok {
		return *stackFrame
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
