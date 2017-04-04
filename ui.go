// Package gocli : Command line interface
//
// These codes are licensed under CC0.
// http://creativecommons.org/publicdomain/zero/1.0/
package gocli

import (
	"bufio"
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

//self-referential function for functional options pattern
type option func(*UI)

// NewUI returns a new UI instance
func NewUI(opts ...option) *UI {
	c := &UI{reader: ioutil.NopCloser(bytes.NewReader(nil)), writer: ioutil.Discard, errorWriter: ioutil.Discard}
	c.Option(opts...)
	return c
}

//Reader returns closure as type option
func Reader(r io.Reader) option {
	return func(c *UI) {
		if r != nil {
			c.reader = r
		}

	}
}

//Writer returns closure as type option
func Writer(w io.Writer) option {
	return func(c *UI) {
		if w != nil {
			c.writer = w
		}
	}
}

//ErrorWriter returns closure as type option
func ErrorWriter(e io.Writer) option {
	return func(c *UI) {
		if e != nil {
			c.errorWriter = e
		}
	}
}

//Option sets options to UI.
func (c *UI) Option(opts ...option) {
	for _, opt := range opts {
		opt(c)
	}
}

//Reader returns io.Reader stream
func (c *UI) Reader() io.Reader {
	return c.reader
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
	writer := bufio.NewWriter(c.writer)
	if _, err := writer.Write(data); err != nil {
		return err
	}
	return writer.Flush()
}

//OutputErr to errorWriter stream.
func (c *UI) OutputErr(val ...interface{}) error {
	return doOutput(c.errorWriter, val)
}

//OutputErrln to errorWriter stream (add line-ending).
func (c *UI) OutputErrln(val ...interface{}) error {
	return doOutputln(c.errorWriter, val)
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
