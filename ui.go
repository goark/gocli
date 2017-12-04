// Package gocli : Command line interface
//
// These codes are licensed under CC0.
// http://creativecommons.org/publicdomain/zero/1.0/
package gocli

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
)

// UI is Command line user interface
type UI struct {
	reader      io.Reader
	writer      io.Writer
	errorWriter io.Writer
}

//OptFunc is self-referential function for functional options pattern
type OptFunc func(*UI)

// NewUI returns a new UI instance
func NewUI(opts ...OptFunc) *UI {
	c := &UI{reader: ioutil.NopCloser(bytes.NewReader(nil)), writer: ioutil.Discard, errorWriter: ioutil.Discard}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

//Reader returns closure as type OptFunc
func Reader(r io.Reader) OptFunc {
	return func(c *UI) {
		if r != nil {
			c.reader = r
		}
	}
}

//Writer returns closure as type OptFunc
func Writer(w io.Writer) OptFunc {
	return func(c *UI) {
		if w != nil {
			c.writer = w
		}
	}
}

//ErrorWriter returns closure as type OptFunc
func ErrorWriter(e io.Writer) OptFunc {
	return func(c *UI) {
		if e != nil {
			c.errorWriter = e
		}
	}
}

//Reader returns io.Reader stream
func (c *UI) Reader() io.Reader {
	return c.reader
}

//Writer returns io.Writer stream for stdout
func (c *UI) Writer() io.Writer {
	return c.writer
}

//ErrorWriter returns io.Writer stream for stderr
func (c *UI) ErrorWriter() io.Writer {
	return c.errorWriter
}

//Output to writer stream.
func (c *UI) Output(val ...interface{}) error {
	return doOutput(c.writer, val)
}

//Outputln to writer stream (add line-ending).
func (c *UI) Outputln(val ...interface{}) error {
	return doOutputln(c.writer, val)
}

//OutputBytes to writer stream ([]byte data).
func (c *UI) OutputBytes(data []byte) error {
	return c.WriteFrom(bytes.NewReader(data))
}

//WriteFrom write from io.Reader to UI.writer
func (c *UI) WriteFrom(r io.Reader) error {
	_, err := io.Copy(c.writer, r)
	return err
}

//OutputErr to errorWriter stream.
func (c *UI) OutputErr(val ...interface{}) error {
	return doOutput(c.errorWriter, val)
}

//OutputErrln to errorWriter stream (add line-ending).
func (c *UI) OutputErrln(val ...interface{}) error {
	return doOutputln(c.errorWriter, val)
}

//OutputErrBytes to writer stream ([]byte data).
func (c *UI) OutputErrBytes(data []byte) error {
	return c.WriteErrFrom(bytes.NewReader(data))
}

//WriteErrFrom write from io.Reader to UI.writer
func (c *UI) WriteErrFrom(r io.Reader) error {
	_, err := io.Copy(c.errorWriter, r)
	return err
}

//Output to io.Writer stream.
func doOutput(writer io.Writer, val []interface{}) error {
	_, err := fmt.Fprint(writer, val...)
	return err
}

//Output to io.Writer stream (add line-ending).
func doOutputln(writer io.Writer, val []interface{}) error {
	_, err := fmt.Fprintln(writer, val...)
	return err
}
