package tools

import (
	"fmt"
	"io/ioutil"
)

const (
	FileMode = 0644
)

// Writer defines the behaviour
// needed to write an array of bytes.
type Writer interface {
	Write(bytes []byte) error
}

type stdWriter struct{}

// NewStdWriter returns a new standard (stdin) writer.
func NewStdWriter() *stdWriter {
	return &stdWriter{}
}

func (*stdWriter) Write(bytes []byte) error {
	_, err := fmt.Println(string(bytes))
	return err
}

type fileWriter struct{ path string }

// NewFileWriter returns a new file writer.
func NewFileWriter(path string) *fileWriter {
	return &fileWriter{path: path}
}

func (fw *fileWriter) Write(bytes []byte) error {
	return ioutil.WriteFile(fw.path, bytes, FileMode)
}
