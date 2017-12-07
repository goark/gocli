// Package rwi : Reader/Writer Interface for command-line
//
// These codes are licensed under CC0.
// http://creativecommons.org/publicdomain/zero/1.0/
package rwi

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
)

// RWI is Reader/Writer class for command-line
type RWI struct {
	reader      io.Reader
	writer      io.Writer
	errorWriter io.Writer
}

//OptFunc is self-referential function for functional options pattern
type OptFunc func(*RWI)

// New returns a new RWI instance
func New(opts ...OptFunc) *RWI {
	c := &RWI{reader: ioutil.NopCloser(bytes.NewReader(nil)), writer: ioutil.Discard, errorWriter: ioutil.Discard}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

//Reader returns closure as type OptFunc
func Reader(r io.Reader) OptFunc {
	return func(c *RWI) {
		if r != nil {
			c.reader = r
		}
	}
}

//Writer returns closure as type OptFunc
func Writer(w io.Writer) OptFunc {
	return func(c *RWI) {
		if w != nil {
			c.writer = w
		}
	}
}

//ErrorWriter returns closure as type OptFunc
func ErrorWriter(e io.Writer) OptFunc {
	return func(c *RWI) {
		if e != nil {
			c.errorWriter = e
		}
	}
}

//Reader returns RWI.reader
func (c *RWI) Reader() io.Reader {
	return c.reader
}

//Writer returns RWI.writer
func (c *RWI) Writer() io.Writer {
	return c.writer
}

//ErrorWriter returns RWI.errorWriter
func (c *RWI) ErrorWriter() io.Writer {
	return c.errorWriter
}

//Output output to RWI.writer
func (c *RWI) Output(val ...interface{}) error {
	return doOutput(c.writer, val)
}

//Outputln output to  RWI.writer (add newline).
func (c *RWI) Outputln(val ...interface{}) error {
	return doOutputln(c.writer, val)
}

//OutputBytes to  RWI.writer ([]byte data).
func (c *RWI) OutputBytes(data []byte) error {
	return c.WriteFrom(bytes.NewReader(data))
}

//WriteFrom  copy from io.Reader to RWI.writer
func (c *RWI) WriteFrom(r io.Reader) error {
	_, err := io.Copy(c.writer, r)
	return err
}

//OutputErr output to  RWI.errorWriter
func (c *RWI) OutputErr(val ...interface{}) error {
	return doOutput(c.errorWriter, val)
}

//OutputErrln output to  RWI.errorWriter (add newline).
func (c *RWI) OutputErrln(val ...interface{}) error {
	return doOutputln(c.errorWriter, val)
}

//OutputErrBytes copy to  RWI.errorWriter ([]byte data).
func (c *RWI) OutputErrBytes(data []byte) error {
	return c.WriteErrFrom(bytes.NewReader(data))
}

//WriteErrFrom copy from io.Reader to RWI.errorWriter
func (c *RWI) WriteErrFrom(r io.Reader) error {
	_, err := io.Copy(c.errorWriter, r)
	return err
}

//Output to io.Writer (internal)
func doOutput(writer io.Writer, val []interface{}) error {
	_, err := fmt.Fprint(writer, val...)
	return err
}

//Output to io.Writer (add newline, internal)
func doOutputln(writer io.Writer, val []interface{}) error {
	_, err := fmt.Fprintln(writer, val...)
	return err
}
